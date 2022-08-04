package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/mitchellnel/lottery-chain-nel/testutil/network"
	"github.com/mitchellnel/lottery-chain-nel/testutil/nullify"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/client/cli"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
)

func networkWithEntranceFeeObjects(t *testing.T) (*network.Network, types.EntranceFee) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	entranceFee := &types.EntranceFee{}
	nullify.Fill(&entranceFee)
	state.EntranceFee = entranceFee
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.EntranceFee
}

func TestShowEntranceFee(t *testing.T) {
	net, obj := networkWithEntranceFeeObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.EntranceFee
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowEntranceFee(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetEntranceFeeResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.EntranceFee)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.EntranceFee),
				)
			}
		})
	}
}
