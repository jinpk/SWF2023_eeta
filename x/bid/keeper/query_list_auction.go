package keeper

import (
	"context"

	"eeta/x/bid/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAuction(goCtx context.Context, req *types.QueryListAuctionRequest) (*types.QueryListAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var auctions []types.Auction

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.GetAuctionKeyPrefix(req.BillboardId))

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var bid types.Auction
		if err := k.cdc.Unmarshal(value, &bid); err != nil {
			return err
		}

		auctions = append(auctions, bid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryListAuctionResponse{
		Pagination: pageRes,
		Auctions:   auctions,
	}, nil
}
