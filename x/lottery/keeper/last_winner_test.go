package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/mitchellnel/lottery-chain-nel/testutil/keeper"
	"github.com/mitchellnel/lottery-chain-nel/testutil/nullify"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/keeper"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
)

func createTestLastWinner(keeper *keeper.Keeper, ctx sdk.Context) types.LastWinner {
	item := types.LastWinner{}
	keeper.SetLastWinner(ctx, item)
	return item
}

func TestLastWinnerGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestLastWinner(keeper, ctx)
	rst, found := keeper.GetLastWinner(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLastWinnerRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestLastWinner(keeper, ctx)
	keeper.RemoveLastWinner(ctx)
	_, found := keeper.GetLastWinner(ctx)
	require.False(t, found)
}
