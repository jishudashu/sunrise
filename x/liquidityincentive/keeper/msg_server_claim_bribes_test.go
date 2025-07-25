package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/sunriselayer/sunrise/x/liquidityincentive/keeper"
	"github.com/sunriselayer/sunrise/x/liquidityincentive/types"
	liquiditypooltypes "github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

func setupEpochs(ctx sdk.Context, k *keeper.Keeper) error {
	// Set up initial epoch
	_ = k.SetEpochCount(ctx, 1)
	initialEpoch := types.Epoch{
		Id:         1,
		StartBlock: ctx.BlockHeight() - 100,
		EndBlock:   ctx.BlockHeight(),
		Gauges: []types.Gauge{{
			PoolId:      1,
			VotingPower: math.NewInt(1),
		}},
	}
	if err := k.SetEpoch(ctx, initialEpoch); err != nil {
		return err
	}

	// Set up current epoch
	currentEpoch := types.Epoch{
		Id:         2,
		StartBlock: ctx.BlockHeight(),
		EndBlock:   ctx.BlockHeight() + 100,
		Gauges: []types.Gauge{{
			PoolId:      1,
			VotingPower: math.NewInt(1),
		}},
	}
	if err := k.SetEpoch(ctx, currentEpoch); err != nil {
		return err
	}

	// Set up next epoch
	nextEpoch := types.Epoch{
		Id:         3,
		StartBlock: ctx.BlockHeight() + 101,
		EndBlock:   ctx.BlockHeight() + 200,
		Gauges: []types.Gauge{{
			PoolId:      1,
			VotingPower: math.NewInt(1),
		}},
	}
	if err := k.SetEpoch(ctx, nextEpoch); err != nil {
		return err
	}

	// Set the current epoch ID to 3
	return k.SetEpochCount(ctx, 3)
}

func TestClaimBribes(t *testing.T) {
	_, _, addr1 := testdata.KeyTestPubAddr()
	addr1Str := addr1.String()
	_, _, addr2 := testdata.KeyTestPubAddr()
	addr2Str := addr2.String()

	bech32Codec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())

	tests := []struct {
		name      string
		msg       func(bribeId uint64) *types.MsgClaimBribes
		expectErr bool
		setup     func(fx *fixture, ctx sdk.Context, bribeId uint64, bribeAmount sdk.Coins)
	}{
		{
			name: "valid bribe claim",
			msg: func(bribeId uint64) *types.MsgClaimBribes {
				return &types.MsgClaimBribes{Sender: addr2Str, BribeIds: []uint64{bribeId}}
			},
			expectErr: false,
			setup: func(fx *fixture, ctx sdk.Context, bribeId uint64, bribeAmount sdk.Coins) {
				fx.mocks.BankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.BribeAccount, addr2, bribeAmount).Return(nil)
			},
		},
		{
			name: "claim non-existent bribe",
			msg: func(_ uint64) *types.MsgClaimBribes {
				return &types.MsgClaimBribes{Sender: addr2Str, BribeIds: []uint64{999}}
			},
			expectErr: true,
		},
		{
			name: "claim with wrong address",
			msg: func(bribeId uint64) *types.MsgClaimBribes {
				return &types.MsgClaimBribes{Sender: addr1Str, BribeIds: []uint64{bribeId}}
			},
			expectErr: true,
		},
		{
			name: "claim already claimed bribe",
			msg: func(bribeId uint64) *types.MsgClaimBribes {
				return &types.MsgClaimBribes{Sender: addr2Str, BribeIds: []uint64{bribeId}}
			},
			expectErr: true,
			setup: func(fx *fixture, ctx sdk.Context, bribeId uint64, bribeAmount sdk.Coins) {
				fx.mocks.BankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.BribeAccount, addr2, bribeAmount).Return(nil)
				msgServer := keeper.NewMsgServerImpl(fx.keeper)
				_, err := msgServer.ClaimBribes(ctx, &types.MsgClaimBribes{Sender: addr2Str, BribeIds: []uint64{bribeId}})
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			fx := initFixture(t)
			sdkCtx := fx.ctx.(sdk.Context)
			msgServer := keeper.NewMsgServerImpl(fx.keeper)

			// Set up mocks for bribe registration first
			fx.mocks.AcctKeeper.EXPECT().AddressCodec().Return(bech32Codec).AnyTimes()
			fx.mocks.LiquiditypoolKeeper.EXPECT().GetPool(gomock.Any(), uint64(1)).Return(liquiditypooltypes.Pool{}, true, nil).AnyTimes()
			fx.mocks.BankKeeper.EXPECT().IsSendEnabledCoins(gomock.Any(), sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(100)))).Return(nil).AnyTimes()
			fx.mocks.BankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), types.BribeAccount, sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(100)))).Return(nil).AnyTimes()

			// Now set up epochs and bribe count
			sdkCtx = sdkCtx.WithBlockHeight(1000)
			err := setupEpochs(sdkCtx, &fx.keeper)
			require.NoError(t, err)
			err = fx.keeper.SetBribeCount(sdkCtx, 1)
			require.NoError(t, err)

			// Set up vote and allocation
			vote := types.Vote{Sender: addr2Str, PoolWeights: []types.PoolWeight{{PoolId: 1, Weight: "1.0"}}}
			err = fx.keeper.SetVote(sdkCtx, vote)
			require.NoError(t, err)
			err = fx.keeper.SaveVoteWeightsForBribes(sdkCtx, 4)
			require.NoError(t, err)
			allocation := types.BribeAllocation{Address: addr2Str, EpochId: 4, PoolId: 1, Weight: "1.0", ClaimedBribeIds: []uint64{}}
			err = fx.keeper.SetBribeAllocation(sdkCtx, allocation)
			require.NoError(t, err)

			// Register a bribe
			bribeAmount := sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(100)))
			sdkCtx = sdkCtx.WithEventManager(sdk.NewEventManager())
			msg := &types.MsgRegisterBribe{Sender: addr1Str, EpochId: 4, PoolId: 1, Amount: bribeAmount}
			_, err = msgServer.RegisterBribe(sdkCtx, msg)
			require.NoError(t, err)
			bribes, err := fx.keeper.GetAllBribeByEpochId(sdkCtx, 4)
			require.NoError(t, err)
			require.Len(t, bribes, 1)
			bribeId := bribes[0].Id
			require.NotZero(t, bribeId)

			sdkCtx = sdkCtx.WithEventManager(sdk.NewEventManager())
			if tc.setup != nil {
				tc.setup(fx, sdkCtx, bribeId, bribeAmount)
			}
			_, err = msgServer.ClaimBribes(sdkCtx, tc.msg(bribeId))
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				bribe, found, err := fx.keeper.GetBribe(sdkCtx, bribeId)
				require.NoError(t, err)
				require.True(t, found)
				require.Equal(t, bribeAmount, bribe.ClaimedAmount)
				allocation, err := fx.keeper.GetBribeAllocation(sdkCtx, addr2, 4, 1)
				require.NoError(t, err)
				require.Contains(t, allocation.ClaimedBribeIds, bribeId)
			}
		})
	}
}
