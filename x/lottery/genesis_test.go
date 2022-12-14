package lottery_test

import (
	"testing"

	keepertest "github.com/mitchellnel/lottery-chain-nel/testutil/keeper"
	"github.com/mitchellnel/lottery-chain-nel/testutil/nullify"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Owner:        &types.Owner{},
		EntranceFee:  &types.EntranceFee{},
		LotteryState: &types.LotteryState{},
		PlayerList: []types.Player{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PlayerCount: 2,
		LastWinner:  &types.LastWinner{},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Owner, got.Owner)
	require.Equal(t, genesisState.EntranceFee, got.EntranceFee)
	require.Equal(t, genesisState.LotteryState, got.LotteryState)
	require.ElementsMatch(t, genesisState.PlayerList, got.PlayerList)
	require.Equal(t, genesisState.PlayerCount, got.PlayerCount)
	require.Equal(t, genesisState.LastWinner, got.LastWinner)
	// this line is used by starport scaffolding # genesis/test/assert
}
