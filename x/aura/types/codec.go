package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBurn{}, "aura/Burn", nil)
	cdc.RegisterConcrete(&MsgMint{}, "aura/Mint", nil)

	cdc.RegisterConcrete(&MsgPause{}, "aura/Pause", nil)
	cdc.RegisterConcrete(&MsgUnpause{}, "aura/Unpause", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgBurn{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgMint{})

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgPause{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUnpause{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
