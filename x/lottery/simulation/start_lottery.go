package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"lottery-chain-nel/x/lottery/keeper"
	"lottery-chain-nel/x/lottery/types"
)

func SimulateMsgStartLottery(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgStartLottery{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the StartLottery simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "StartLottery simulation not implemented"), nil, nil
	}
}
