package keeper

import (
	"context"

	"eeta/x/deposit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Coin)); err != nil {
		return nil, err
	}

	k.Logger(ctx).Info(types.ModuleName)
	k.Logger(ctx).Info(msg.ReceipientAddress)
	k.Logger(ctx).Info(sdk.NewCoins(msg.Coin).String())

	if address, err := sdk.AccAddressFromBech32(msg.ReceipientAddress); err != nil {
		return nil, err
	} else {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(msg.Coin)); err != nil {
			return nil, err
		}
	}

	return &types.MsgMintResponse{}, nil
}
