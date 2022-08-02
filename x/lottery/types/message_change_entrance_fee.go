package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChangeEntranceFee = "change_entrance_fee"

var _ sdk.Msg = &MsgChangeEntranceFee{}

func NewMsgChangeEntranceFee(creator string, newEntranceFee uint64) *MsgChangeEntranceFee {
	return &MsgChangeEntranceFee{
		Creator:        creator,
		NewEntranceFee: newEntranceFee,
	}
}

func (msg *MsgChangeEntranceFee) Route() string {
	return RouterKey
}

func (msg *MsgChangeEntranceFee) Type() string {
	return TypeMsgChangeEntranceFee
}

func (msg *MsgChangeEntranceFee) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeEntranceFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeEntranceFee) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
