package keeper

import (
	"context"

	"eeta/x/deposit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 권한체크
	minterAddress := sdk.MustAccAddressFromBech32(msg.Sender)
	if err := k.assertMinter(ctx, minterAddress); err != nil {
		return nil, types.ErrUnauthMint
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Coin)); err != nil {
		return nil, err
	}

	if address, err := sdk.AccAddressFromBech32(msg.ReceipientAddress); err != nil {
		return nil, err
	} else {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.NewCoins(msg.Coin)); err != nil {
			return nil, err
		}
	}

	return &types.MsgMintResponse{}, nil
}
