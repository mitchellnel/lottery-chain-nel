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

	// MARK: there is no good documentation on how to write integration tests for Cosmos SDK-built
	// apps

	// Was given https://github.com/cosmos/b9-checkers-academy-draft/tree/v1-main as an example for
	// tests, but it has so many additional files for the testing, and I don't know what they do,
	// so right now, I'm ignoring integration testing for the enter-lottery message

	// Basically, my issue is that enter-lottery also interacts with the bank module, and I have no
	// way to simulate the bank module right now

	// I'll just test using the command line when the chain is served
}
