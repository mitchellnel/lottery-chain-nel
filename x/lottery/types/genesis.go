package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Owner:        nil,
		EntranceFee:  nil,
		LotteryState: nil,
		PlayerList:   []Player{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in player
	playerIndexMap := make(map[string]struct{})

	for _, elem := range gs.PlayerList {
		index := string(PlayerKey(elem.Address))
		if _, ok := playerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for player")
		}
		playerIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
