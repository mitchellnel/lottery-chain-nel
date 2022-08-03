package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"lottery-chain-nel/x/lottery/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated player",
			genState: &types.GenesisState{
				PlayerList: []types.Player{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid player count",
			genState: &types.GenesisState{
				PlayerList: []types.Player{
					{
						Id: 1,
					},
				},
				PlayerCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
