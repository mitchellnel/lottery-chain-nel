package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lottery-chain-nel/x/lottery/types"
)

func (k Keeper) EntranceFee(c context.Context, req *types.QueryGetEntranceFeeRequest) (*types.QueryGetEntranceFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEntranceFee(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEntranceFeeResponse{EntranceFee: val}, nil
}
