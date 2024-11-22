package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"
	keepertest "github.com/sunriselayer/sunrise/testutil/keeper"
	"github.com/sunriselayer/sunrise/testutil/nullify"
	"github.com/sunriselayer/sunrise/x/liquiditypool/keeper"
	"github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

func createNPosition(keeper keeper.Keeper, ctx context.Context, n int) []types.Position {
	items := make([]types.Position, n)
	for i := range items {
		items[i] = types.Position{
			Id:        0,
			Address:   fmt.Sprintf("%d", i),
			PoolId:    1,
			LowerTick: 1,
			UpperTick: 2,
			Liquidity: math.LegacyOneDec(),
		}
		items[i].Id = keeper.AppendPosition(ctx, items[i])
	}
	return items
}

func TestPositionGet(t *testing.T) {
	keeper, _, ctx := keepertest.LiquiditypoolKeeper(t)
	items := createNPosition(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPosition(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPositionRemove(t *testing.T) {
	keeper, _, ctx := keepertest.LiquiditypoolKeeper(t)
	items := createNPosition(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePosition(ctx, item.Id)
		_, found := keeper.GetPosition(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPositionGetAll(t *testing.T) {
	keeper, _, ctx := keepertest.LiquiditypoolKeeper(t)
	items := createNPosition(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPositions(ctx)),
	)
}

func TestPositionCount(t *testing.T) {
	keeper, _, ctx := keepertest.LiquiditypoolKeeper(t)
	items := createNPosition(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPositionCount(ctx))
}
