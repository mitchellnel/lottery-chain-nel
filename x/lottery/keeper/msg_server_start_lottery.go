package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery-chain-nel/x/lottery/types"
)

func (k msgServer) StartLottery(goCtx context.Context, msg *types.MsgStartLottery) (*types.MsgStartLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgStartLotteryResponse{}, nil
}
