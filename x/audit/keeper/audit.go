package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/mattverse/errata/x/audit/types"
)

func (k Keeper) RegisterProtocol(ctx sdk.Context, title string, description string, sourceCode string, projectHome string, category string) error {
	id := k.GetLastProtocolID(ctx) + 1
	protocol := types.Protocol{
		Id:          id,
		Title:       title,
		Description: description,
		SourceCode:  sourceCode,
		ProjectHome: projectHome,
		Category:    category,
	}
	err := k.setProtocol(ctx, protocol)
	if err != nil {
		return err
	}
	return nil
}

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

func (k Keeper) SetProtocols(ctx sdk.Context, protocols []types.Protocol) error {
	for _, protocol := range protocols {
		err := k.setProtocol(ctx, protocol)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) GetLastProtocolID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyLastProtocolID)
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetLastProtocolID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyLastProtocolID, sdk.Uint64ToBigEndian(id))
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

func (k Keeper) AddAttackPoolByProtocolID(ctx sdk.Context, protocolID uint64, amount sdk.Int) error {
	protocol, err := k.GetProtocolByID(ctx, protocolID)
	if err != nil {
		return err
	}

	protocol.AttackPool = protocol.AttackPool.Add(amount)
	err = k.setProtocol(ctx, *protocol)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) AddDefensePoolByProtocolID(ctx sdk.Context, protocolID uint64, amount sdk.Int) error {
	protocol, err := k.GetProtocolByID(ctx, protocolID)
	if err != nil {
		return err
	}

	protocol.DefensePool = protocol.DefensePool.Add(amount)
	err = k.setProtocol(ctx, *protocol)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) AddErrata(ctx sdk.Context, protocolID uint64, vulnerabilityType string, errataCode string, vulnerability string) error {
	protocol, err := k.GetProtocolByID(ctx, protocolID)
	if err != nil {
		return err
	}

	errata := types.Errata{
		Vulnerability:     vulnerability,
		ErrataCode:        errataCode,
		VulnerabilityType: vulnerabilityType,
	}

	protocol.Errata = append(protocol.Errata, &errata)
	err = k.setProtocol(ctx, *protocol)

	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) ProtocolIterator(ctx sdk.Context) sdk.Iterator {
	return k.iterator(ctx, types.KeyProtocol)
}

func (k Keeper) iterator(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}
