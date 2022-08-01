package lottery

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"lottery-chain-nel/testutil/sample"
	lotterysimulation "lottery-chain-nel/x/lottery/simulation"
	"lottery-chain-nel/x/lottery/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = lotterysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgClaimOwner = "op_weight_msg_claim_owner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimOwner int = 100

	opWeightMsgChangeOwner = "op_weight_msg_change_owner"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChangeOwner int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	lotteryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&lotteryGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgClaimOwner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaimOwner, &weightMsgClaimOwner, nil,
		func(_ *rand.Rand) {
			weightMsgClaimOwner = defaultWeightMsgClaimOwner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimOwner,
		lotterysimulation.SimulateMsgClaimOwner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgChangeOwner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgChangeOwner, &weightMsgChangeOwner, nil,
		func(_ *rand.Rand) {
			weightMsgChangeOwner = defaultWeightMsgChangeOwner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChangeOwner,
		lotterysimulation.SimulateMsgChangeOwner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
