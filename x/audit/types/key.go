package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// KeyProtocol defines state to store protocol
	KeyProtocol = []byte{0x01}

	// KeyIndexSeparator defines separator between keys when combine, it should be one that is not used in denom expression
	KeyIndexSeparator = []byte{0xFF}
)

func ProtocolStoreKey(ID uint64) []byte {
	return combineKeys(KeyProtocol, sdk.Uint64ToBigEndian(ID))
}

// combineKeys combine bytes array into a single bytes
func combineKeys(keys ...[]byte) []byte {
	return bytes.Join(keys, KeyIndexSeparator)
}
