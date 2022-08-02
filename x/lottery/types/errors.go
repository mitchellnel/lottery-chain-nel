package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lottery module sentinel errors
var (
	ErrLotteryAlreadyOwned = sdkerrors.Register(ModuleName, 1000, "Lottery already has an owner")
	ErrLotteryNotOwned     = sdkerrors.Register(ModuleName, 1001, "Lottery does not have an owner")

	ErrNotOwner = sdkerrors.Register(ModuleName, 2000, "You are not the lottery owner")

	ErrLotteryNotSetup = sdkerrors.Register(
		ModuleName,
		3000,
		"Lottery has not been setup before",
	)
	ErrLotteryAlreadySetup = sdkerrors.Register(
		ModuleName,
		3001,
		"Lottery has already been setup before",
	)

	ErrLotteryNotClosed = sdkerrors.Register(
		ModuleName,
		4000,
		"Lottery is not closed",
	)
	ErrLotteryClosed = sdkerrors.Register(
		ModuleName,
		4001,
		"Lottery is closed",
	)
)
