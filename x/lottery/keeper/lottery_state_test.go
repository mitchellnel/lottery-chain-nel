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

func createTestLotteryState(keeper *keeper.Keeper, ctx sdk.Context) types.LotteryState {
	item := types.LotteryState{}
	keeper.SetLotteryState(ctx, item)
	return item
}

func createTestClosedLotteryState(keeper *keeper.Keeper, ctx sdk.Context) types.LotteryState {
	item := types.LotteryState{LotteryState: types.LotteryState_CLOSED}
	keeper.SetLotteryState(ctx, item)
	return item
}

func createTestOpenLotteryState(keeper *keeper.Keeper, ctx sdk.Context) types.LotteryState {
	item := types.LotteryState{LotteryState: types.LotteryState_OPEN}
	keeper.SetLotteryState(ctx, item)
	return item
}

func TestLotteryStateGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestLotteryState(keeper, ctx)
	rst, found := keeper.GetLotteryState(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLotteryStateRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestLotteryState(keeper, ctx)
	keeper.RemoveLotteryState(ctx)
	_, found := keeper.GetLotteryState(ctx)
	require.False(t, found)
}
