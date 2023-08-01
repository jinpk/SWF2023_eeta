package keeper

import (
	"context"

	"eeta/x/bid/types"

	deposittypes "eeta/x/deposit/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

	// 낙찰 희망 가격 잔고 확인
	has := k.bk.GetBalance(ctx, creatorAddr, deposittypes.StableCoinDenom)
	if has.Amount.LT(msg.Amount.Amount) {
		return nil, types.ErrInsufientBalance
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAuctionKeyPrefix(msg.BillboardId))

	// 겹치는 시간대에 진행중인 비딩이 있는지 확인
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var auction types.Auction
		k.cdc.MustUnmarshal(iterator.Value(), &auction)
		// 절대시간 검증 로직 없음 단순 계산
		// TODO: 겹치는 시간 검증 필요
		if msg.Start >= auction.Start && msg.End <= auction.End ||
			msg.Start >= auction.Start && msg.Start+msg.End > auction.End ||
			msg.Start >= auction.Start && auction.Start+auction.End > msg.End ||
			msg.End <= auction.End && msg.Start+msg.End > auction.Start {
			return nil, types.ErrExistAuctionTime
		}
	}

	var nextAuctionId uint64 = k.NextAuctionId(ctx, msg.BillboardId)

	// auction 생성
	auction := types.Auction{
		Id:            nextAuctionId,
		BillboardId:   msg.BillboardId,
		Start:         msg.Start,
		End:           msg.End,
		BidderAddress: "", // 낙찰자 없음
	}
	bz := k.cdc.MustMarshal(&auction)
	store.Set(types.GetIDBytes(nextAuctionId), bz)

	// bid 추가
	bid := types.Bid{
		SenderAddress: msg.Creator,
		Amount:        msg.Amount,
		AdUrl:         msg.AdUrl,
	}

	bidBz := k.cdc.MustMarshal(&bid)

	auctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetBidKeyPrefix(msg.BillboardId, auction.Id))

	auctionStore.Set(creatorAddr, bidBz)

	return &types.MsgCreateAuctionResponse{}, nil
}
