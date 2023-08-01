package keeper

import (
	"context"

	"eeta/x/deposit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	burnedAddress := sdk.MustAccAddressFromBech32(msg.BurnedAddress)

	// 권한체크
	minterAddress := sdk.MustAccAddressFromBech32(msg.Creator)
	if err := k.assertMinter(ctx, minterAddress); err != nil {
		return nil, types.ErrUnauthMint
	}

	// 잔고부족
	if !k.bankKeeper.GetBalance(ctx, burnedAddress, msg.Amount.Denom).Amount.GTE(msg.Amount.Amount) {
		return nil, types.ErrInsufientBurnBalance
	}

	// address to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, burnedAddress, types.ModuleName, sdk.NewCoins(msg.Amount)); err != nil {
		return nil, types.ErrFailedBurn
	}

	// burn
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Amount)); err != nil {
		return nil, types.ErrFailedBurn
	}

	return &types.MsgBurnResponse{}, nil
}
