package keeper

import (
	"context"

	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) StartLottery(
	goCtx context.Context,
	msg *types.MsgStartLottery,
) (*types.MsgStartLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that the message creator is the owner of the lottery
	owner, found := k.GetOwner(ctx)
	if !found || (owner.Owner != msg.Creator) {
		return &types.MsgStartLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrNotOwner,
				"You are not the lottery owner, and therefore cannot start the lottery",
			)
	}

	// check that the lottery has undergone its first-time setup
	lotteryState, found := k.GetLotteryState(ctx)
	if !found {
		return &types.MsgStartLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryNotSetup,
				"Lottery has not yet been setup -- send a setup-lottery message to set it up",
			)
	}

	// check that the lottery has not already been started (i.e. it is open)
	if lotteryState.LotteryState == types.LotteryState_OPEN {
		return &types.MsgStartLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryOpen,
				"Lottery is already open",
			)
	}

	// clear the list of Players
	k.RemoveAllPlayer(ctx)

	// change the lottery state to OPEN
	var newLotteryState = types.LotteryState{
		LotteryState: types.LotteryState_OPEN,
	}
	k.SetLotteryState(ctx, newLotteryState)

	return &types.MsgStartLotteryResponse{Success: true}, nil
}
