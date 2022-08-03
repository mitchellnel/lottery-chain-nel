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
	cdc.RegisterConcrete(&MsgChangeEntranceFee{}, "lottery/ChangeEntranceFee", nil)
	cdc.RegisterConcrete(&MsgEnterLottery{}, "lottery/EnterLottery", nil)
	cdc.RegisterConcrete(&MsgStartLottery{}, "lottery/StartLottery", nil)
	cdc.RegisterConcrete(&MsgEndLottery{}, "lottery/EndLottery", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeEntranceFee{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnterLottery{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStartLottery{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEndLottery{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
