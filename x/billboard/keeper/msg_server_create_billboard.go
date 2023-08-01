package keeper

import (
	"context"

	"eeta/x/billboard/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBillboard(goCtx context.Context, msg *types.MsgCreateBillboard) (*types.MsgCreateBillboardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BillboardKey))

	billboardId := k.GetNextBillboardId(ctx)

	billboard := types.Billboard{
		Id:           billboardId,
		OwnerAddress: msg.Creator,
		Name:         msg.Name,
		Description:  msg.Description,
		Url:          msg.Url,
		BoardType:    msg.BoardType,
		MinPrice:     msg.MinPrice,
	}

	bz := k.cdc.MustMarshal(&billboard)
	store.Set(types.GetBillboardIDBytes(billboard.Id), bz)

	return &types.MsgCreateBillboardResponse{}, nil
}
