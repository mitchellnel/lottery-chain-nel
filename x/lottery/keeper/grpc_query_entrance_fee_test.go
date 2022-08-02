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

func TestEntranceFeeQuery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestEntranceFee(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetEntranceFeeRequest
		response *types.QueryGetEntranceFeeResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetEntranceFeeRequest{},
			response: &types.QueryGetEntranceFeeResponse{EntranceFee: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EntranceFee(wctx, tc.request)
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
