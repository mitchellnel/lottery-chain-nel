package keeper_test

import (
	keepertest "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/x/lottery/keeper"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestEnterLottery(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgServer := keeper.NewMsgServerImpl(*k)

	_, _, _, _ = k, ctx, wctx, msgServer
}
