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

func createTestEntranceFee(keeper *keeper.Keeper, ctx sdk.Context) types.EntranceFee {
	item := types.EntranceFee{}
	keeper.SetEntranceFee(ctx, item)
	return item
}

func createTestEntranceFeeWithValue(keeper *keeper.Keeper, ctx sdk.Context) types.EntranceFee {
	item := types.EntranceFee{EntranceFee: 44}
	keeper.SetEntranceFee(ctx, item)
	return item
}

func TestEntranceFeeGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestEntranceFee(keeper, ctx)
	rst, found := keeper.GetEntranceFee(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestEntranceFeeRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestEntranceFee(keeper, ctx)
	keeper.RemoveEntranceFee(ctx)
	_, found := keeper.GetEntranceFee(ctx)
	require.False(t, found)
}
