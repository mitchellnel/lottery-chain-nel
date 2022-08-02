package keeper_test

import (
	keepertest "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/x/lottery/keeper"
	"lottery-chain-nel/x/lottery/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestChangeEntranceFee(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgServer := keeper.NewMsgServerImpl(*k)

	/* ChangeEntranceFee when message creator is the lottery owner, and the lottery has been setup
	   before*/
	_, _, _ = createNamedTestOwner(
		k,
		ctx,
	), createTestClosedLotteryState(
		k,
		ctx,
	), createTestEntranceFeeWithValue(
		k,
		ctx,
	)

	var message = types.MsgChangeEntranceFee{
		Creator:        "test-owner",
		NewEntranceFee: 3,
	}

	var expectedResponse = types.MsgChangeEntranceFeeResponse{
		Success: true,
	}

	response, err := msgServer.ChangeEntranceFee(wctx, &message)

	// check correct response, and no error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.NoError(t, err)

	// check that entrance fee was changed correctly
	entranceFee, found := k.GetEntranceFee(ctx)
	require.Equal(t, true, found)
	require.Equal(t, uint64(3), entranceFee.EntranceFee)

	///

	/* ChangeEntranceFee when lottery is not owned */
	k.RemoveOwner(ctx)
	k.RemoveLotteryState(ctx)
	k.RemoveEntranceFee(ctx)

	message = types.MsgChangeEntranceFee{
		Creator:        "no-owner",
		NewEntranceFee: 33,
	}

	expectedResponse = types.MsgChangeEntranceFeeResponse{
		Success: false,
	}

	response, err = msgServer.ChangeEntranceFee(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryNotOwned, err)

	///

	/* ChangeEntranceFee when message creator is not the owner of the lottery */
	_, _, expectedEntranceFee := createNamedTestOwner(
		k,
		ctx,
	), createTestClosedLotteryState(
		k,
		ctx,
	), createTestEntranceFeeWithValue(
		k,
		ctx,
	)

	message = types.MsgChangeEntranceFee{
		Creator:        "wrong-owner",
		NewEntranceFee: 33,
	}

	expectedResponse = types.MsgChangeEntranceFeeResponse{
		Success: false,
	}

	response, err = msgServer.ChangeEntranceFee(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrNotOwner, err)

	// check that entrance fee is unchanged
	entranceFee, found = k.GetEntranceFee(ctx)
	require.Equal(t, true, found)
	require.Equal(t, expectedEntranceFee.EntranceFee, entranceFee.EntranceFee)

	///

	/* ChangeEntranceFee when lottery has not been setup before */
	k.RemoveOwner(ctx)
	k.RemoveLotteryState(ctx)
	k.RemoveEntranceFee(ctx)

	_ = createNamedTestOwner(k, ctx)

	message = types.MsgChangeEntranceFee{
		Creator:        "test-owner",
		NewEntranceFee: 33,
	}

	expectedResponse = types.MsgChangeEntranceFeeResponse{
		Success: false,
	}

	response, err = msgServer.ChangeEntranceFee(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryNotSetup, err)

	///

	/* ChangeEntranceFee when the lottery is closed */
	k.RemoveOwner(ctx)
	k.RemoveLotteryState(ctx)
	k.RemoveEntranceFee(ctx)

	_, _, expectedEntranceFee = createNamedTestOwner(
		k,
		ctx,
	), createTestOpenLotteryState(
		k,
		ctx,
	), createTestEntranceFeeWithValue(
		k,
		ctx,
	)

	message = types.MsgChangeEntranceFee{
		Creator:        "test-owner",
		NewEntranceFee: 33,
	}

	expectedResponse = types.MsgChangeEntranceFeeResponse{
		Success: false,
	}

	response, err = msgServer.ChangeEntranceFee(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryNotClosed, err)

	// check that entrance fee is unchanged
	entranceFee, found = k.GetEntranceFee(ctx)
	require.Equal(t, true, found)
	require.Equal(t, expectedEntranceFee.EntranceFee, entranceFee.EntranceFee)
}
