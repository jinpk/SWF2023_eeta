package keeper

import (
	"context"
	"strings"

	"eeta/x/bid/types"
	deposittypes "eeta/x/deposit/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Bidding(goCtx context.Context, msg *types.MsgBidding) (*types.MsgBiddingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

	auctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAuctionKeyPrefix(msg.BillboardId))

	auctionBz := auctionStore.Get(types.GetIDBytes(msg.AuctionId))

	// 옥션 있는지 확인
	if auctionBz == nil {
		return nil, types.ErrNotFoundAuction
	}

	// 낙찰 희망 가격 잔고 확인
	has := k.bk.GetBalance(ctx, creatorAddr, deposittypes.StableCoinDenom)
	if has.Amount.LT(msg.Amount.Amount) {
		return nil, types.ErrInsufientBalance
	}

	var auction types.Auction
	k.cdc.MustUnmarshal(auctionBz, &auction)

	// 이미 옥션 끝났는지 확인
	if !strings.EqualFold(auction.BidderAddress, "") {
		return nil, types.ErrAuctionBidded
	}

	// TODO: 기존 입찰 이력 있다면 환수 후 재 입찰

	bidStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetBidKeyPrefix(msg.BillboardId, msg.AuctionId))

	if bidStore.Has(creatorAddr) {
		return nil, types.ErrAlreadyBidded
	}

	// bid 추가
	bid := types.Bid{
		SenderAddress: msg.Creator,
		Amount:        msg.Amount,
		AdUrl:         msg.AdUrl,
	}
	bidBz := k.cdc.MustMarshal(&bid)

	bidStore.Set(creatorAddr, bidBz)

	return &types.MsgBiddingResponse{}, nil
}
