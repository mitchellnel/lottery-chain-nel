package lottery_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lottery-chain-nel/testutil/keeper"
	"lottery-chain-nel/testutil/nullify"
	"lottery-chain-nel/x/lottery"
	"lottery-chain-nel/x/lottery/types"
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
	// this line is used by starport scaffolding # genesis/test/assert
}
