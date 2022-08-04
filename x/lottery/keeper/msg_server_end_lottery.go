package keeper

import (
	"context"
	"math/rand"
	"time"

	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) EndLottery(
	goCtx context.Context,
	msg *types.MsgEndLottery,
) (*types.MsgEndLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that the message creator is the owner of the lottery
	owner, found := k.GetOwner(ctx)
	if !found || (owner.Owner != msg.Creator) {
		return &types.MsgEndLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrNotOwner,
				"You are not the lottery owner, and therefore cannot end the lottery",
			)
	}

	// check that the lottery is currently open
	lotteryState, _ := k.GetLotteryState(ctx)
	if lotteryState.LotteryState != types.LotteryState_OPEN {
		return &types.MsgEndLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryNotOpen,
				"The lottery is not open, and therefore cannot be closed",
			)
	}

	// deal with edge case where there were no entered players
	// in this case, we have no funds to dispense and no winner to choose
	// so just close the lottery without error
	if k.GetPlayerCount(ctx) == 0 {
		var closedLotteryState = types.LotteryState{
			LotteryState: types.LotteryState_CLOSED,
		}

		k.SetLotteryState(ctx, closedLotteryState)

		return &types.MsgEndLotteryResponse{Success: true}, nil
	}

	// MARK: typically, we should use some oracle with a VRF for true, verifiable randomness --
	// however, BandChain, the supported oracle for Cosmos SDK-built chains does not yet offer VRF,
	// so Go's random module will just have to suffice since this chain isn't for production use

	// close the lottery
	// MARK: if we had a VRF, we would label it as CALCULATING_WINNER
	// (maybe we should still do this to indicate we're in this block of code, something to
	// consider)
	var closedLotteryState = types.LotteryState{
		LotteryState: types.LotteryState_CLOSED,
	}
	k.SetLotteryState(ctx, closedLotteryState)

	// get the lottery players
	players := k.GetAllPlayer(ctx)
	nPlayers := len(players)

	// set a seed to try and produce non-deterministic random numbers
	rand.Seed(time.Now().UnixNano())

	// choose a random player ID in the range [0, nPlayers) == [0, nPlayers-1]
	winnerId := rand.Intn(nPlayers)

	// get the winner and get their address
	winner := players[winnerId]
	winnerAddr, err := sdk.AccAddressFromBech32(winner.Address)
	if err != nil {
		return &types.MsgEndLotteryResponse{Success: false}, err
	}

	// update last-winner
	var lastWinner = types.LastWinner{
		Address: winner.Address,
	}
	k.SetLastWinner(ctx, lastWinner)

	// get address of the lottery module account
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))

	// transfer the winnings (i.e. all the funds in the module) to the winner
	moduleBalance := k.bankKeeper.GetBalance(ctx, moduleAcct, "token")

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		winnerAddr,
		sdk.NewCoins(moduleBalance),
	)
	if err != nil {
		return &types.MsgEndLotteryResponse{Success: false}, err
	}

	return &types.MsgEndLotteryResponse{Success: true}, nil
}
