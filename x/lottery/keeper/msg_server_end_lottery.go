package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery-chain-nel/x/lottery/types"
)

func (k msgServer) EndLottery(goCtx context.Context, msg *types.MsgEndLottery) (*types.MsgEndLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEndLotteryResponse{}, nil
}
