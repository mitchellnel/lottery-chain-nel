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

func createTestOwner(keeper *keeper.Keeper, ctx sdk.Context) types.Owner {
	item := types.Owner{}
	keeper.SetOwner(ctx, item)
	return item
}

func createNamedTestOwner(keeper *keeper.Keeper, ctx sdk.Context) types.Owner {
	item := types.Owner{Owner: "test-owner"}
	keeper.SetOwner(ctx, item)
	return item
}

func TestOwnerGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestOwner(keeper, ctx)
	rst, found := keeper.GetOwner(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestOwnerRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestOwner(keeper, ctx)
	keeper.RemoveOwner(ctx)
	_, found := keeper.GetOwner(ctx)
	require.False(t, found)
}
