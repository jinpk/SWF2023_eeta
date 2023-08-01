package keeper

import (
	"context"

	"eeta/x/sto/types"

	billboardtypes "eeta/x/billboard/types"
	deposittypes "eeta/x/deposit/types"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSto(goCtx context.Context, msg *types.MsgCreateSto) (*types.MsgCreateStoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.bk.Has(ctx, msg.BillboardId) {
		return nil, billboardtypes.ErrNotFoundBillboard
	}

	if msg.UserShare+msg.OrganizerShare != 100 {
		return nil, types.ErrInvalidShareTotal
	}

	duraionMin := (msg.End - msg.Start) / 60
	pricePerMinute := k.bk.GetFinalBidPricePerMinute(ctx, msg.BillboardId)
	goal := sdk.NewCoin(pricePerMinute.Denom, pricePerMinute.Amount.Mul(math.NewInt(duraionMin)))

	nextStoId := k.NextStoId(ctx, msg.BillboardId)
	creatorAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	stoAddresss := types.NewStoPoolAddress(nextStoId)

	// start, end 시간 등록 가능한지 계산

	// TODO: creator shares로 펀드 계산 (나누기)
	organizerFundAmount := sdk.NewCoin(deposittypes.StableCoinDenom, math.NewInt(1000_000000))

	if k.bankKeeper.GetBalance(ctx, creatorAddr, deposittypes.StableCoinDenom).Amount.LT(organizerFundAmount.Amount) {
		return nil, types.ErrInsufiendOrganazierShare
	}

	sto := types.Sto{
		Id:                nextStoId,
		BillboardId:       msg.BillboardId,
		OrganizerAddress:  msg.Creator,
		OrganizerUrl:      msg.OrganizerUrl,
		OrganizerImageUrl: msg.OrganizerImageUrl,
		Name:              msg.Name,
		Start:             msg.Start,
		End:               msg.End,
		UserShare:         msg.UserShare,
		OrganizerShare:    msg.OrganizerShare,
		Goal:              goal,
		Funded:            organizerFundAmount,
		StoAddress:        stoAddresss.String(),
	}

	// sto 저장
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetStoKeyPrefix(msg.BillboardId))
	store.Set(types.GetIDBytes(sto.Id), k.cdc.MustMarshal(&sto))

	// 코인 펀딩 deposit
	if err := k.bankKeeper.SendCoins(ctx, creatorAddr, stoAddresss, sdk.NewCoins(organizerFundAmount)); err != nil {
		return nil, types.ErrFailedStoFunding
	}

	// sto asset 발행
	denom := types.GetStoTokenDenom(nextStoId)
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, goal.Amount)))

	// orginaizer 에게 sto 발행
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddr, sdk.NewCoins(sdk.NewCoin(denom, organizerFundAmount.Amount)))

	return &types.MsgCreateStoResponse{}, nil
}
