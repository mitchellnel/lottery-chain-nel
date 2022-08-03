package keeper_test

import (
	"testing"

	keepertest "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/testutil/nullify"
	"lottery-chain-nel/x/lottery/keeper"
	"lottery-chain-nel/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNPlayer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Player {
	items := make([]types.Player, n)
	for i := range items {
		items[i].Id = keeper.AppendPlayer(ctx, items[i])
	}
	return items
}

func TestPlayerGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNPlayer(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPlayer(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPlayerRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNPlayer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePlayer(ctx, item.Id)
		_, found := keeper.GetPlayer(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPlayerGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNPlayer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPlayer(ctx)),
	)
}

func TestPlayerRemoveAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	_ = createNPlayer(keeper, ctx, 10)

	keeper.RemoveAllPlayer(ctx)

	require.Equal(t, uint64(0), keeper.GetPlayerCount(ctx))
	require.Empty(t, keeper.GetAllPlayer(ctx))
}

func TestPlayerCount(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNPlayer(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPlayerCount(ctx))
}
