package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LastWinner(
	c context.Context,
	req *types.QueryGetLastWinnerRequest,
) (*types.QueryGetLastWinnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLastWinner(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLastWinnerResponse{LastWinner: val}, nil
}
