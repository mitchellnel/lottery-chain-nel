package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartLottery = "start_lottery"

var _ sdk.Msg = &MsgStartLottery{}

func NewMsgStartLottery(creator string) *MsgStartLottery {
	return &MsgStartLottery{
		Creator: creator,
	}
}

func (msg *MsgStartLottery) Route() string {
	return RouterKey
}

func (msg *MsgStartLottery) Type() string {
	return TypeMsgStartLottery
}

func (msg *MsgStartLottery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartLottery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
