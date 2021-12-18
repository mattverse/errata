package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgRegisterProtocol = "register_protocol"
	TypeMsgJoinAttackPool   = "join_attack_pool"
	TypeMsgJoinDefensePool  = "join_defense_pool"
	TypeMsgAddErrata        = "add_errata"
)

var (
	_ sdk.Msg = &MsgRegisterProtocol{}
	_ sdk.Msg = &MsgJoinAttackPool{}
	_ sdk.Msg = &MsgJoinDefensePool{}
	_ sdk.Msg = &MsgAddErrata{}
)

func NewMsgRegisterProtocol(
	title string,
	description string,
	sourceCode string,
	projectHome string,
	category string,
	sender sdk.AccAddress) *MsgRegisterProtocol {
	return &MsgRegisterProtocol{
		Title:       title,
		Description: description,
		SourceCode:  sourceCode,
		ProjectHome: projectHome,
		Category:    category,
		Sender:      sender.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgRegisterProtocol) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgRegisterProtocol) Type() string { return TypeMsgRegisterProtocol }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgRegisterProtocol) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{sender}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgRegisterProtocol) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgRegisterProtocol) ValidateBasic() error {
	return nil
}

func NewMsgJoinAttackPool(poolID uint64, tokenIn sdk.Int, sender sdk.AccAddress) *MsgJoinAttackPool {
	return &MsgJoinAttackPool{
		PoolId:  poolID,
		TokenIn: tokenIn,
		Sender:  sender.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgJoinAttackPool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgJoinAttackPool) Type() string { return TypeMsgRegisterProtocol }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgJoinAttackPool) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{sender}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgJoinAttackPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgJoinAttackPool) ValidateBasic() error {
	return nil
}

func NewMsgJoinDefensePool(poolID uint64, tokenIn sdk.Int, sender sdk.AccAddress) *MsgJoinDefensePool {
	return &MsgJoinDefensePool{
		PoolId:  poolID,
		TokenIn: tokenIn,
		Sender:  sender.String(),
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgJoinDefensePool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgJoinDefensePool) Type() string { return TypeMsgRegisterProtocol }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgJoinDefensePool) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{sender}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgJoinDefensePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgJoinDefensePool) ValidateBasic() error {
	return nil
}

func NewMsgAddErrata(poolID uint64, vulnerabilityType string, errataCode string, vulnerability string) *MsgAddErrata {
	return &MsgAddErrata{
		PoolId:            poolID,
		VulnerabilityType: vulnerabilityType,
		ErrataCode:        errataCode,
		Vulnerability:     vulnerability,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgAddErrata) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgAddErrata) Type() string { return TypeMsgRegisterProtocol }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgAddErrata) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{sender}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgAddErrata) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgAddErrata) ValidateBasic() error {
	return nil
}
