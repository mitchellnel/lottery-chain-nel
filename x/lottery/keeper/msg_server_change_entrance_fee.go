package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery-chain-nel/x/lottery/types"
)

func (k msgServer) ChangeEntranceFee(goCtx context.Context, msg *types.MsgChangeEntranceFee) (*types.MsgChangeEntranceFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgChangeEntranceFeeResponse{}, nil
}
