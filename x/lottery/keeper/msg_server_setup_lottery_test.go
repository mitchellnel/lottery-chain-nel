package keeper_test

import (
	keepertest "github.com/mitchellnel/lottery-chain-nel/testutil/keeper"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/keeper"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestSetupLottery(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgServer := keeper.NewMsgServerImpl(*k)

	/* SetupLottery when message creator is the lottery owner, and the lottery has not been setup
	   before */
	_ = createNamedTestOwner(k, ctx)

	var message = types.MsgSetupLottery{
		Creator:     "test-owner",
		EntranceFee: 3,
	}

	var expectedResponse = types.MsgSetupLotteryResponse{
		Success: true,
	}

	response, err := msgServer.SetupLottery(wctx, &message)

	// check correct response, and no error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.NoError(t, err)

	// check that the lottery was setup correctly, and entrance fee is set
	entranceFee, found := k.GetEntranceFee(ctx)
	require.Equal(t, true, found)
	require.Equal(t, uint64(3), entranceFee.EntranceFee)

	lotteryState, found := k.GetLotteryState(ctx)
	require.Equal(t, true, found)
	require.Equal(t, types.LotteryState_CLOSED, lotteryState.LotteryState)

	///

	/* SetupLottery when lottery is not owned */
	k.RemoveOwner(ctx)
	k.RemoveEntranceFee(ctx)
	k.RemoveLotteryState(ctx)

	message = types.MsgSetupLottery{
		Creator:     "no-owner",
		EntranceFee: 33,
	}

	expectedResponse = types.MsgSetupLotteryResponse{
		Success: false,
	}

	response, err = msgServer.SetupLottery(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryNotOwned, err)

	// check that entrance fee is not set
	_, found = k.GetEntranceFee(ctx)
	require.Equal(t, false, found)

	// check that lottery state is not set
	_, found = k.GetLotteryState(ctx)
	require.Equal(t, false, found)

	///

	/* SetupLottery when message creator is not the owner of the lottery */
	_ = createNamedTestOwner(k, ctx)

	message = types.MsgSetupLottery{
		Creator:     "wrong-owner",
		EntranceFee: 33,
	}

	expectedResponse = types.MsgSetupLotteryResponse{
		Success: false,
	}

	response, err = msgServer.SetupLottery(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrNotOwner, err)

	// check that entrance fee is not set
	_, found = k.GetEntranceFee(ctx)
	require.Equal(t, false, found)

	// check that lottery state is not set
	_, found = k.GetLotteryState(ctx)
	require.Equal(t, false, found)

	///

	/* SetupLottery when lottery has been set up before */
	// if lottery has been setup before, it will have an entrance fee
	expectedEntranceFee := createTestEntranceFee(k, ctx)
	expectedLotteryState := createTestLotteryState(k, ctx)

	message = types.MsgSetupLottery{
		Creator:     "test-owner",
		EntranceFee: 33,
	}

	expectedResponse = types.MsgSetupLotteryResponse{
		Success: false,
	}

	response, err = msgServer.SetupLottery(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryAlreadySetup, err)

	// check that entrance fee is still in its original state
	entranceFee, _ = k.GetEntranceFee(ctx)
	require.Equal(t, expectedEntranceFee, entranceFee)

	// check that lottery state is still in its original state
	lotteryState, _ = k.GetLotteryState(ctx)
	require.Equal(t, expectedLotteryState, lotteryState)
}
