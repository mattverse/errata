package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mattverse/errata/x/audit/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the bond MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) RegisterProtocol(goCtx context.Context, msg *types.MsgRegisterProtocol) (*types.MsgRegisterProtocolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.keeper.RegisterProtocol(ctx, msg.Title, msg.Description, msg.SourceCode, msg.ProjectHome, msg.Category)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterProtocolResponse{}, nil
}

func (k msgServer) JoinAttackPool(goCtx context.Context, msg *types.MsgJoinAttackPool) (*types.MsgJoinAttackPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.keeper.AddAttackPoolByProtocolId(ctx, msg.PoolId, msg.TokenIn)
	if err != nil {
		return nil, err
	}

	return &types.MsgJoinAttackPoolResponse{}, nil
}

func (k msgServer) JoinDefensePool(goCtx context.Context, msg *types.MsgJoinDefensePool) (*types.MsgJoinDefensePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.keeper.AddDefensePoolByProtocolId(ctx, msg.PoolId, msg.TokenIn)
	if err != nil {
		return nil, err
	}

	return &types.MsgJoinDefensePoolResponse{}, nil
}

func (k msgServer) AddErrata(goCtx context.Context, msg *types.MsgAddErrata) (*types.MsgAddErrataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.keeper.AddErrata(ctx, msg.PoolId, msg.Vulnerability, msg.ErrataCode, msg.Vulnerability)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddErrataResponse{}, nil
}
