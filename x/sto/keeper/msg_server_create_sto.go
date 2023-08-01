package keeper

import (
	"context"

	"eeta/x/sto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSto(goCtx context.Context, msg *types.MsgCreateSto) (*types.MsgCreateStoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateStoResponse{}, nil
}
