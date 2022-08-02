package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery-chain-nel/x/lottery/keeper"
	"lottery-chain-nel/x/lottery/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.Owner != nil {
		k.SetOwner(ctx, *genState.Owner)
	}
	// Set if defined
	if genState.EntranceFee != nil {
		k.SetEntranceFee(ctx, *genState.EntranceFee)
	}
	// Set if defined
	if genState.LotteryState != nil {
		k.SetLotteryState(ctx, *genState.LotteryState)
	}
	// Set all the player
	for _, elem := range genState.PlayerList {
		k.SetPlayer(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all owner
	owner, found := k.GetOwner(ctx)
	if found {
		genesis.Owner = &owner
	}
	// Get all entranceFee
	entranceFee, found := k.GetEntranceFee(ctx)
	if found {
		genesis.EntranceFee = &entranceFee
	}
	// Get all lotteryState
	lotteryState, found := k.GetLotteryState(ctx)
	if found {
		genesis.LotteryState = &lotteryState
	}
	genesis.PlayerList = k.GetAllPlayer(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
