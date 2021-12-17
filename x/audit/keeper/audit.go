package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/mattverse/errata/x/audit/types"
)

func (k Keeper) GetProtocolByID(ctx sdk.Context, protocolID uint64) (*types.Protocol, error) {
	protocol := types.Protocol{}
	store := ctx.KVStore(k.storeKey)
	protocolKey := types.ProtocolStoreKey(protocolID)
	if !store.Has(protocolKey) {
		return nil, fmt.Errorf("protocol with ID %d does not exist", protocolID)
	}
	bz := store.Get(protocolKey)
	err := proto.Unmarshal(bz, &protocol)
	return &protocol, err
}

func (k Keeper) GetAllProtocol(ctx sdk.Context) []types.Protocol {
	protocols := []types.Protocol{}
	iterator := k.ProtocolIterator(ctx)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		protocolID := sdk.BigEndianToUint64(iterator.Value())
		protocol, err := k.GetProtocolByID(ctx, protocolID)
		if err != nil {
			panic(err)
		}
		protocols = append(protocols, *protocol)
	}
	return protocols

}

func (k Keeper) setProtocol(ctx sdk.Context, protocol types.Protocol) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := proto.Marshal(&protocol)
	if err != nil {
		return err
	}
	store.Set(types.ProtocolStoreKey(protocol.Id), bz)
	return nil
}

func (k Keeper) AddAttackPoolByProtocolId(ctx sdk.Context, protocolID uint64, amount sdk.Int) error {
	protocol, err := k.GetProtocolByID(ctx, protocolID)
	if err != nil {
		return err
	}

	protocol.AttackPool = protocol.AttackPool.Add(amount)
	k.setProtocol(ctx, *protocol)
	return nil
}

func (k Keeper) AddDefensePoolByProtocolId(ctx sdk.Context, protocolID uint64, amount sdk.Int) error {
	protocol, err := k.GetProtocolByID(ctx, protocolID)
	if err != nil {
		return err
	}

	protocol.DefensePool = protocol.DefensePool.Add(amount)
	k.setProtocol(ctx, *protocol)
	return nil
}

func (k Keeper) AddErrata(ctx sdk.Context, protocolID uint64, errata types.Errata) error {
	protocol, err := k.GetProtocolByID(ctx, protocolID)
	if err != nil {
		return err
	}

	protocol.Errata = append(protocol.Errata, &errata)
	k.setProtocol(ctx, *protocol)
	return nil
}

func (k Keeper) ProtocolIterator(ctx sdk.Context) sdk.Iterator {
	return k.iterator(ctx, types.KeyProtocol)
}

func (k Keeper) iterator(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}
