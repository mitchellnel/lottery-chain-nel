package keeper

import (
	"context"

	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ChangeEntranceFee(
	goCtx context.Context,
	msg *types.MsgChangeEntranceFee,
) (*types.MsgChangeEntranceFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that the lottery is owned
	currentOwner, alreadyOwned := k.GetOwner(ctx)
	if !alreadyOwned {
		return &types.MsgChangeEntranceFeeResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryNotOwned,
				"Lottery does not have an owner yet -- claim ownership first",
			)
	}

	// check that the message creator is the current lottery owner
	if currentOwner.Owner != msg.Creator {
		return &types.MsgChangeEntranceFeeResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrNotOwner,
				"Message creator does not own the lottery",
			)
	}

	// check that the lottery has been set up before
	lotteryState, setupBefore := k.GetLotteryState(ctx)
	if !setupBefore {
		return &types.MsgChangeEntranceFeeResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryNotSetup,
				"Lottery has not yet been set up -- set up lottery first",
			)
	}

	// check that the lottery is currently closed
	if lotteryState.LotteryState != types.LotteryState_CLOSED {
		return &types.MsgChangeEntranceFeeResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryNotClosed,
				"Lottery is not closed -- lottery must be closed to change entrance fee",
			)
	}

	// change the entrance fee to the new entrance fee
	var newEntranceFee = types.EntranceFee{
		EntranceFee: msg.NewEntranceFee,
	}

	k.SetEntranceFee(ctx, newEntranceFee)

	return &types.MsgChangeEntranceFeeResponse{Success: true}, nil
}
