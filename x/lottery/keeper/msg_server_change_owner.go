package keeper

import (
	"context"

	"lottery-chain-nel/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ChangeOwner(
	goCtx context.Context,
	msg *types.MsgChangeOwner,
) (*types.MsgChangeOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that the lottery is owned
	currentOwner, alreadyOwned := k.GetOwner(ctx)
	if !alreadyOwned {
		return &types.MsgChangeOwnerResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryNotOwned,
				"Lottery does not have an owner yet -- claim ownership first",
			)
	}

	// check that the message creator is the current lottery owner
	if currentOwner.Owner != msg.Creator {
		return &types.MsgChangeOwnerResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrNotOwner,
				"Message creator does not own the lottery",
			)
	}

	// change owner to the new owner
	var newOwner = types.Owner{
		Owner: msg.NewOwner,
	}

	k.SetOwner(ctx, newOwner)

	return &types.MsgChangeOwnerResponse{Success: true}, nil
}
