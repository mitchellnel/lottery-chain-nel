package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEndLottery = "end_lottery"

var _ sdk.Msg = &MsgEndLottery{}

func NewMsgEndLottery(creator string) *MsgEndLottery {
	return &MsgEndLottery{
		Creator: creator,
	}
}

func (msg *MsgEndLottery) Route() string {
	return RouterKey
}

func (msg *MsgEndLottery) Type() string {
	return TypeMsgEndLottery
}

func (msg *MsgEndLottery) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEndLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEndLottery) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
