package keeper

import (
	"context"
	"fmt"
	"time"

	"eeta/x/bid/types"

	billboardtypes "eeta/x/billboard/types"
	deposittypes "eeta/x/deposit/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr := sdk.MustAccAddressFromBech32(msg.Creator)

	now := time.Now()
	startTime := time.Unix(int64(msg.Start), 0)

	// 이전 시간 생성 불가
	if startTime.Before(now) {
		return nil, types.ErrBeforeTime
	}

	params := k.GetParams(ctx)
	blockGenSecond := 10 // TODO: 블록 생성 시간 계산 필요
	// 공백 사간 필요
	if now.Add(time.Second * time.Duration(int(params.WaitBlock)*blockGenSecond)).After(startTime) {
		return nil, types.ErrRequireSpaceTime
	}

	// 낙찰 희망 가격 잔고 확인
	has := k.bk.GetBalance(ctx, creatorAddr, deposittypes.StableCoinDenom)
	if has.Amount.LT(msg.Amount.Amount) {
		return nil, types.ErrInsufientBalance
	}

	if !k.dk.Has(ctx, msg.BillboardId) {
		return nil, billboardtypes.ErrNotFoundBillboard
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAuctionKeyPrefix(msg.BillboardId))

	// 겹치는 시간대에 진행중인 비딩이 있는지 확인
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var auction types.Auction
		k.cdc.MustUnmarshal(iterator.Value(), &auction)
		if msg.Start >= auction.Start && msg.End <= auction.End {
			return nil, types.ErrExistAuctionTime
		}
		//if msg.Start >= auction.Start && msg.End > auction.End {
		//	return nil, types.ErrExistAuctionTime
		//}
		//if msg.End <= auction.End && msg.Start+msg.End > auction.Start {
		//	return nil, types.ErrExistAuctionTime
		//}
	}

	var nextAuctionId uint64 = k.NextAuctionId(ctx, msg.BillboardId)

	// auction 생성
	auctionPoolAddr := types.NewAuctionPoolAddress(nextAuctionId)
	k.Logger(ctx).Info(fmt.Sprintf("new auction address %d %s", nextAuctionId, auctionPoolAddr.String()))
	auction := types.Auction{
		AuctionAddress: auctionPoolAddr.String(),
		Id:             nextAuctionId,
		BillboardId:    msg.BillboardId,
		Start:          msg.Start,
		End:            msg.End,
		BidderAddress:  "", // 낙찰자 없음
	}
	bz := k.cdc.MustMarshal(&auction)
	store.Set(types.GetIDBytes(nextAuctionId), bz)

	// bid 추가
	bid := types.Bid{
		SenderAddress: msg.Creator,
		Amount:        msg.Amount,
		AdUrl:         msg.AdUrl,
		Height:        uint64(ctx.BlockHeight()),
	}

	bidBz := k.cdc.MustMarshal(&bid)

	auctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetBidKeyPrefix(msg.BillboardId, auction.Id))

	if err := k.bk.SendCoins(ctx, creatorAddr, auctionPoolAddr, sdk.NewCoins(bid.Amount)); err != nil {
		return nil, types.ErrDepositFailed
	}

	k.Logger(ctx).Info("deposit account to bid module for bidding")

	auctionStore.Set(creatorAddr, bidBz)

	return &types.MsgCreateAuctionResponse{}, nil
}
