package keeper

import (
	"context"

	"eeta/x/sto/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Fund(goCtx context.Context, msg *types.MsgFund) (*types.MsgFundResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetStoKeyPrefix(msg.BillboardId))
	stoBz := store.Get(types.GetIDBytes(msg.StoId))
	var sto types.Sto
	k.cdc.MustUnmarshal(stoBz, &sto)

	creatorAddr := sdk.MustAccAddressFromBech32(msg.Creator)

	// TODO: start, end에 맞춰 펀드 가입 가능한지 검증 필요

	remain := sto.Goal.Sub(sto.Funded)
	if remain.IsZero() {
		return nil, types.ErrAllFunded
	}

	var fund sdk.Coin = msg.Amount
	if msg.Amount.IsGTE(remain) {
		fund = remain
	}

	sto.Funded = sto.Funded.Add(fund)

	store.Set(types.GetIDBytes(sto.Id), k.cdc.MustMarshal(&sto))

	// 코인 펀딩 deposit
	if err := k.bankKeeper.SendCoins(ctx, creatorAddr, sdk.MustAccAddressFromBech32(sto.StoAddress), sdk.NewCoins(fund)); err != nil {
		return nil, types.ErrFailedStoFunding
	}

	// sto asset 발행
	denom := types.GetStoTokenDenom(sto.Id)
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, fund.Amount)))

	// orginaizer 에게 sto 발행
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddr, sdk.NewCoins(sdk.NewCoin(denom, fund.Amount)))

	return &types.MsgFundResponse{}, nil
}
