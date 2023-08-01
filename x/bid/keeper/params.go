package keeper

import (
	"eeta/x/bid/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.WaitBlock(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// WaitBlock returns the WaitBlock param
func (k Keeper) WaitBlock(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyWaitBlock, &res)
	return
}
