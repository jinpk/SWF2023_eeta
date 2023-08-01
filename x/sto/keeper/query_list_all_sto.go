package keeper

import (
	"context"

	"eeta/x/sto/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAllSto(goCtx context.Context, req *types.QueryListAllStoRequest) (*types.QueryListAllStoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stos []types.Sto
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, []byte(types.StoKey))

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var sto types.Sto
		if err := k.cdc.Unmarshal(value, &sto); err != nil {
			return err
		}

		stos = append(stos, sto)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListAllStoResponse{
		Pagination: pageRes,
		Stos:       stos,
	}, nil
}
