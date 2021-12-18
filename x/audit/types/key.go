package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the audit module
	ModuleName = "audit"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the staking module
	RouterKey = ModuleName

	QuerierRoute = ModuleName
)

var (
	// KeyProtocol defines state to store protocol
	KeyProtocol = []byte{0x01}

	// KeyLastProtocolID defines key to store protocol ID used by last
	KeyLastProtocolID = []byte{0x02}

	// KeyIndexSeparator defines separator between keys when combine, it should be one that is not used in denom expression
	KeyIndexSeparator = []byte{0xFF}
)

func ProtocolStoreKey(id uint64) []byte {
	return combineKeys(KeyProtocol, sdk.Uint64ToBigEndian(id))
}

// combineKeys combine bytes array into a single bytes
func combineKeys(keys ...[]byte) []byte {
	return bytes.Join(keys, KeyIndexSeparator)
}
