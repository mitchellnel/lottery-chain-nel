package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgClaimOwner{}, "lottery/ClaimOwner", nil)
	cdc.RegisterConcrete(&MsgChangeOwner{}, "lottery/ChangeOwner", nil)
	cdc.RegisterConcrete(&MsgSetupLottery{}, "lottery/SetupLottery", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimOwner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeOwner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetupLottery{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
