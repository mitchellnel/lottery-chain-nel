package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetupLottery = "setup_lottery"

var _ sdk.Msg = &MsgSetupLottery{}

func NewMsgSetupLottery(creator string, entranceFee uint64) *MsgSetupLottery {
	return &MsgSetupLottery{
		Creator:     creator,
		EntranceFee: entranceFee,
	}
}

func (msg *MsgSetupLottery) Route() string {
	return RouterKey
}

func (msg *MsgSetupLottery) Type() string {
	return TypeMsgSetupLottery
}

func (msg *MsgSetupLottery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetupLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetupLottery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
