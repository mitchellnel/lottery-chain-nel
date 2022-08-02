package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/testutil/nullify"
	"lottery-chain-nel/x/lottery/types"
)

func TestLotteryStateQuery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestLotteryState(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLotteryStateRequest
		response *types.QueryGetLotteryStateResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLotteryStateRequest{},
			response: &types.QueryGetLotteryStateResponse{LotteryState: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LotteryState(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
