package keeper

import (
	"context"

	"lottery-chain-nel/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnterLottery(
	goCtx context.Context,
	msg *types.MsgEnterLottery,
) (*types.MsgEnterLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that the lottery has been set up
	lotteryState, stateFound := k.GetLotteryState(ctx)
	if !stateFound {
		return &types.MsgEnterLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryNotSetup,
				"Lottery is not yet setup -- ask the owner to setup the lottery",
			)
	}

	// check that the lottery is open (i.e. not closed)
	if lotteryState.LotteryState != types.LotteryState_OPEN {
		return &types.MsgEnterLotteryResponse{Success: false},
			sdkerrors.Wrapf(
				types.ErrLotteryClosed,
				"Lottery is closed -- wait for the lottery to start and be open",
			)
	}

	// prepare to transfer
	wrapped_entranceFee, _ := k.GetEntranceFee(ctx)
	entranceFee := wrapped_entranceFee.EntranceFee
	entranceFee_coin := sdk.Coins{sdk.NewInt64Coin("token", int64(entranceFee))}

	playerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgEnterLotteryResponse{Success: false}, err
	}

	// tranfer entrance fee from the message creator to the module
	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		playerAddr,
		types.ModuleName,
		entranceFee_coin,
	)
	if err != nil {
		return &types.MsgEnterLotteryResponse{Success: false}, err
	}

	// TODO: add player to a list of players currently taking part in the lottery

	return &types.MsgEnterLotteryResponse{}, nil
}
