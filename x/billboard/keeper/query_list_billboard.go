package keeper

import (
	"context"

	"eeta/x/billboard/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListBillboard(goCtx context.Context, req *types.QueryListBillboardRequest) (*types.QueryListBillboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var billboards []types.Billboard
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.KeyPrefix(types.BillboardKey))

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var billboard types.Billboard
		if err := k.cdc.Unmarshal(value, &billboard); err != nil {
			return err
		}

		billboards = append(billboards, billboard)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListBillboardResponse{
		Pagination: pageRes,
		Billboard:  billboards,
	}, nil
}
