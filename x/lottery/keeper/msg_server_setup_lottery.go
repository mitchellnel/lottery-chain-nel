package keeper

import (
	"context"

	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetupLottery(
	goCtx context.Context,
	msg *types.MsgSetupLottery,
) (*types.MsgSetupLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that the lottery has an owner
	owner, ownerFound := k.GetOwner(ctx)
	if !ownerFound {
		return &types.MsgSetupLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryNotOwned,
				"Lottery does not have an owner yet -- claim ownership first",
			)
	}

	// check that message sender is the lottery owner
	if msg.Creator != owner.Owner {
		return &types.MsgSetupLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrNotOwner,
				"Message creator does not own the lottery",
			)
	}

	// check that lottery has not been set up before
	_, stateFound := k.GetLotteryState(ctx)
	if stateFound {
		return &types.MsgSetupLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryAlreadySetup,
				"Lottery has already been setup",
			)
	}

	// create and set the entrance fee
	var entranceFee = types.EntranceFee{
		EntranceFee: msg.EntranceFee,
	}
	k.SetEntranceFee(ctx, entranceFee)

	// create and set the lottery state
	var lotteryState = types.LotteryState{
		LotteryState: types.LotteryState_CLOSED,
	}
	k.SetLotteryState(ctx, lotteryState)

	return &types.MsgSetupLotteryResponse{Success: true}, nil
}
