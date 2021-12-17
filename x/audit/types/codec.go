package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/stake interfaces and
// concrete types on the provided LegacyAmino codec. These types are used for
// Amino JSON serialization.
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterProtocol{}, "errata/audit/MsgRegisterProtocol", nil)
	cdc.RegisterConcrete(&MsgJoinAttackPool{}, "errata/audit/MsgJoinAttackPool", nil)
	cdc.RegisterConcrete(&MsgJoinDefensePool{}, "errata/audit/MsgJoinDefensePool", nil)
	cdc.RegisterConcrete(&MsgAddErrata{}, "errata/audit/MsgAddErrata", nil)

}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterProtocol{},
		&MsgJoinAttackPool{},
		&MsgJoinDefensePool{},
		&MsgAddErrata{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
