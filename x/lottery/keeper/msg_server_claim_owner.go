package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery-chain-nel/x/lottery/types"
)

func (k msgServer) ClaimOwner(goCtx context.Context, msg *types.MsgClaimOwner) (*types.MsgClaimOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgClaimOwnerResponse{}, nil
}
