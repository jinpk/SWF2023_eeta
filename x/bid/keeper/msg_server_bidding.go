package keeper

import (
	"context"

	"eeta/x/bid/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Bidding(goCtx context.Context, msg *types.MsgBidding) (*types.MsgBiddingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBiddingResponse{}, nil
}
