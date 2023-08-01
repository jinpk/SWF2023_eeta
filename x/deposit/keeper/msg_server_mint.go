package keeper

import (
	"context"

	"eeta/x/deposit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: 스테이블 코인 발행 제한
	//params := k.GetParams(ctx)
	//minterAddres := params.GetMinterAddress()
	//if !strings.EqualFold(minterAddres, msg.Sender) {
	//	return nil, types.ErrUnauthMint
	//}

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
