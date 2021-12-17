package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/mattverse/errata/x/audit/types"
)

// keeper of the bond store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryMarshaler
	bankKeeper types.BankKeeper
}

// NewKeeper creates a new bond Keeper instance
func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, bk types.BankKeeper, ps paramtypes.Subspace) Keeper {
	// set KeyTable if it has not already been set

	return Keeper{
		storeKey:   key,
		cdc:        cdc,
		bankKeeper: bk,
	}
}
