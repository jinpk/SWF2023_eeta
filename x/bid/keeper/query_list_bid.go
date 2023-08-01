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

func (k Keeper) ListBid(goCtx context.Context, req *types.QueryListBidRequest) (*types.QueryListBidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var bids []types.Bid

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.GetBidKeyPrefix(req.BillboardId, req.AuctionId))

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var bid types.Bid
		if err := k.cdc.Unmarshal(value, &bid); err != nil {
			return err
		}

		bids = append(bids, bid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListBidResponse{
		Pagination: pageRes,
		Bids:       bids,
	}, nil
}
