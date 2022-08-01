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
)
