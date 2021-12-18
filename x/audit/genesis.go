package audit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mattverse/errata/x/audit/keeper"
	"github.com/mattverse/errata/x/audit/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.SetProtocols(ctx, genState.Protocols); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {

	genesis := types.DefaultGenesis()
	genesis.Protocols, _ = k.GetAllProtocol(ctx)

	return genesis
}
