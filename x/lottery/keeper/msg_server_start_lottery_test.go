package keeper_test

import (
	keepertest "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/x/lottery/keeper"
	"lottery-chain-nel/x/lottery/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestStartLottery(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgServer := keeper.NewMsgServerImpl(*k)

	_ = createNamedTestOwner(k, ctx)

	var setupMsg = types.MsgSetupLottery{
		Creator:     "test-owner",
		EntranceFee: 3,
	}
	_, err := msgServer.SetupLottery(wctx, &setupMsg)
	require.NoError(t, err)

	/* StartLottery when message creator is the lottery owner, and the lottery has been setup and
	   is currently close */
	var message = types.MsgStartLottery{
		Creator: "test-owner",
	}

	var expectedResponse = types.MsgStartLotteryResponse{
		Success: true,
	}

	response, err := msgServer.StartLottery(wctx, &message)

	// check correct response, and no error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.NoError(t, err)

	// check that lottery is open
	lotteryState, _ := k.GetLotteryState(ctx)
	require.Equal(t, types.LotteryState_OPEN, lotteryState.LotteryState)

	/* StartLottery when the message creator is not the lottery owner */
	k.RemoveEntranceFee(ctx)
	k.RemoveLotteryState(ctx)
	_, err = msgServer.SetupLottery(wctx, &setupMsg)
	require.NoError(t, err)

	message = types.MsgStartLottery{
		Creator: "wrong-owner",
	}

	expectedResponse = types.MsgStartLotteryResponse{
		Success: false,
	}

	response, err = msgServer.StartLottery(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrNotOwner, err)

	// check that lottery state is still closed
	lotteryState, _ = k.GetLotteryState(ctx)
	require.Equal(t, types.LotteryState_CLOSED, lotteryState.LotteryState)

	/* StartLottery when the lottery has not yet been setup */
	k.RemoveEntranceFee(ctx)
	k.RemoveLotteryState(ctx)

	message = types.MsgStartLottery{
		Creator: "test-owner",
	}

	expectedResponse = types.MsgStartLotteryResponse{
		Success: false,
	}

	response, err = msgServer.StartLottery(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryNotSetup, err)

	/* StartLottery when the lottery is already open */
	_, err = msgServer.SetupLottery(wctx, &setupMsg)
	require.NoError(t, err)

	var openState = types.LotteryState{
		LotteryState: types.LotteryState_OPEN,
	}
	k.SetLotteryState(ctx, openState)

	message = types.MsgStartLottery{
		Creator: "test-owner",
	}

	expectedResponse = types.MsgStartLotteryResponse{
		Success: false,
	}

	response, err = msgServer.StartLottery(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryOpen, err)
}
