package bid

import (
	"eeta/x/bid/keeper"
	"eeta/x/bid/types"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)

	waitBlock := params.GetWaitBlock()
	curHeight := ctx.BlockHeight()

	auctions := k.ListUnBiddedAuctions(ctx)

	for _, auction := range auctions {
		bids := k.ListBids(ctx, auction.BillboardId, auction.Id)

		var maxNum uint64
		var bidded *types.Bid
		bidded = &bids[0]
		maxNum = bidded.Height

		for _, bid := range bids {
			bidded = &bid
			maxNum = uint64(math.Max(float64(maxNum), float64(bidded.Height)))
		}

		if uint64(curHeight) > bidded.Height+waitBlock {
			// 낙찰 처리
			k.Bid(ctx, auction.BillboardId, auction.Id, bidded.SenderAddress)
		}
	}
}
