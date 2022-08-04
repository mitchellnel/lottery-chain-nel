package keeper_test

import (
	keepertest "github.com/mitchellnel/lottery-chain-nel/testutil/keeper"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/keeper"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestChangeOwner(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgServer := keeper.NewMsgServerImpl(*k)

	/* ChangeOwner when message creator is the lottery owner */
	_ = createNamedTestOwner(k, ctx)

	var message = types.MsgChangeOwner{
		Creator:  "test-owner",
		NewOwner: "pass",
	}

	var expectedResponse = types.MsgChangeOwnerResponse{
		Success: true,
	}

	response, err := msgServer.ChangeOwner(wctx, &message)

	// check correct response, and no error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.NoError(t, err)

	// check that owner was successfully changed
	owner, found := k.GetOwner(ctx)
	require.Equal(t, true, found)
	require.Equal(t, "pass", owner.Owner)

	///

	/* ChangeOwner when lottery is unowned */
	k.RemoveOwner(ctx)

	message = types.MsgChangeOwner{
		Creator:  "no-owner",
		NewOwner: "fail",
	}

	expectedResponse = types.MsgChangeOwnerResponse{
		Success: false,
	}

	response, err = msgServer.ChangeOwner(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrLotteryNotOwned, err)

	// check that lottery is still unowned
	_, found = k.GetOwner(ctx)
	require.Equal(t, false, found)

	///

	/* ChangeOwner when message creator is not the lottery owner */
	expectedOwner := createNamedTestOwner(k, ctx)

	message = types.MsgChangeOwner{
		Creator:  "wrong-owner",
		NewOwner: "fail",
	}

	expectedResponse = types.MsgChangeOwnerResponse{
		Success: false,
	}

	response, err = msgServer.ChangeOwner(wctx, &message)

	// check correct response, and correct error
	require.Equal(t, expectedResponse.Success, response.Success)
	require.Equal(t, expectedResponse, *response)
	require.ErrorIs(t, types.ErrNotOwner, err)

	// check that lottery is still owned by original owner
	owner, _ = k.GetOwner(ctx)
	require.Equal(t, expectedOwner.Owner, owner.Owner)
}
