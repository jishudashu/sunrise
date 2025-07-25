package keeper

import (
	"context"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/sunriselayer/sunrise/x/liquidityincentive/types"
)

func (k Keeper) BeginBlocker(ctx context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, telemetry.Now(), telemetry.MetricKeyBeginBlocker)

	cacheCtx, write := sdk.UnwrapSDKContext(ctx).CacheContext()

	// Transfer a portion of inflation rewards from fee collector to `x/liquidityincentive` pool.
	feeCollector := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	incentiveModule := authtypes.NewModuleAddress(types.ModuleName)
	tokenconverterParams, err := k.tokenConverterKeeper.GetParams(cacheCtx)
	if err != nil {
		k.Logger().Error("failed to get token converter params", "error", err)
		return nil
	}
	transferableDenom := tokenconverterParams.TransferableDenom
	nonTransferableDenom := tokenconverterParams.NonTransferableDenom

	// Check the Gauge count is not zero.
	// Distribute incentives to gauges
	lastEpoch, found, err := k.GetLastEpoch(cacheCtx)
	if err != nil {
		k.Logger().Error("failed to get last epoch", "error", err)
		return nil
	}
	if !found {
		k.Logger().Info("last epoch not found")
		return nil
	}

	totalPower := math.LegacyZeroDec()
	for _, gauge := range lastEpoch.Gauges {
		totalPower = totalPower.Add(math.LegacyNewDecFromInt(gauge.VotingPower))
	}

	if totalPower.IsZero() {
		k.Logger().Info("total count is zero")
		return nil
	}

	// Send a portion of inflation rewards from fee collector to `x/liquidityincentive` pool.
	feeBalance := k.bankKeeper.GetBalance(cacheCtx, feeCollector, transferableDenom)
	feeCollectorAmountDec := math.LegacyNewDecFromInt(feeBalance.Amount)
	params, err := k.Params.Get(cacheCtx)
	if err != nil {
		k.Logger().Error("failed to get params", "error", err)
		return nil
	}
	stakingRewardRatioDec, err := math.LegacyNewDecFromStr(params.StakingRewardRatio)
	if err != nil {
		k.Logger().Error("failed to parse staking reward ratio", "error", err)
		return nil
	}
	incentiveAmount := feeCollectorAmountDec.Mul(math.LegacyOneDec().Sub(stakingRewardRatioDec)).TruncateInt()
	err = k.bankKeeper.SendCoinsFromModuleToModule(cacheCtx, authtypes.FeeCollectorName, types.ModuleName, sdk.NewCoins(sdk.NewCoin(transferableDenom, incentiveAmount)))
	if err != nil {
		k.Logger().Error("failed to send coins from fee collector to liquidity incentive module", "error", err)
		return nil
	}

	// Convert fee denom to bond denom in the `x/liquidityincentive` module account.
	err = k.tokenConverterKeeper.ConvertReverse(cacheCtx, incentiveAmount, incentiveModule)
	if err != nil {
		k.Logger().Error("failed to convert fee denom to bond denom", "error", err)
		return nil
	}

	// Get `x/liquidityincentive` module's incentive balance.
	incentiveBalance := k.bankKeeper.GetBalance(cacheCtx, incentiveModule, nonTransferableDenom)

	// Distribute incentives to gauges
	for _, gauge := range lastEpoch.Gauges {
		weight := math.LegacyNewDecFromInt(gauge.VotingPower).Quo(totalPower)
		allocationDec := math.LegacyNewDecFromInt(incentiveBalance.Amount).Mul(weight)
		if allocationDec.IsPositive() {
			err := k.liquidityPoolKeeper.AllocateIncentive(
				cacheCtx,
				gauge.PoolId,
				incentiveModule,
				sdk.NewCoins(sdk.NewCoin(nonTransferableDenom, allocationDec.TruncateInt())),
			)
			if err != nil {
				k.Logger().Error("failure in incentive allocation", "error", err)
				return nil
			}
		}
	}

	write()
	return nil
}
