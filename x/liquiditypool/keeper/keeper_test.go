package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/core/address"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectestutil "github.com/cosmos/cosmos-sdk/codec/testutil"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	liquiditypooltestutil "github.com/sunriselayer/sunrise/x/liquiditypool/testutil"
	"go.uber.org/mock/gomock"

	"github.com/sunriselayer/sunrise/x/liquiditypool/keeper"
	module "github.com/sunriselayer/sunrise/x/liquiditypool/module"
	"github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

type fixture struct {
	ctx          context.Context
	keeper       keeper.Keeper
	addressCodec address.Codec
	mocks        LiquidityPoolMocks
}

func initFixture(t *testing.T) *fixture {
	t.Helper()

	config := sdk.GetConfig()
	encCfg := moduletestutil.MakeTestEncodingConfig(codectestutil.CodecOptions{}, module.AppModule{})
	addressCodec := addresscodec.NewBech32Codec(config.GetBech32AccountAddrPrefix())
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	env := runtime.NewEnvironment(runtime.NewKVStoreService(storeKey), log.NewTestLogger(t))
	ctx := testutil.DefaultContextWithDB(t, storeKey, storetypes.NewTransientStoreKey("transient_test")).Ctx

	authority := authtypes.NewModuleAddress(types.GovModuleName)

	mocks := getMocks(t)

	k := keeper.NewKeeper(
		env,
		encCfg.Codec,
		addressCodec,
		authority,
		mocks.BankKeeper,
		nil,
	)

	// Initialize params
	if err := k.Params.Set(ctx, types.DefaultParams()); err != nil {
		t.Fatalf("failed to set params: %v", err)
	}

	return &fixture{
		ctx:          ctx,
		keeper:       k,
		addressCodec: addressCodec,
		mocks:        mocks,
	}
}

type LiquidityPoolMocks struct {
	BankKeeper *liquiditypooltestutil.MockBankKeeper
}

func getMocks(t *testing.T) LiquidityPoolMocks {
	ctrl := gomock.NewController(t)

	return LiquidityPoolMocks{
		BankKeeper: liquiditypooltestutil.NewMockBankKeeper(ctrl),
	}
}
