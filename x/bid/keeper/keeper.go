package keeper

import (
	"fmt"
	"strings"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"eeta/x/bid/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		ak         types.AccountKeeper
		bk         types.BankKeeper
		dk         types.BillboardKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	dk types.BillboardKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		ak:         ak,
		bk:         bk, dk: dk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) NextAuctionId(ctx sdk.Context, billboardId uint64) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAuctionKeyPrefix(billboardId))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var id uint64 = 1

	for ; iterator.Valid(); iterator.Next() {
		id++
	}

	return id
}

func (k Keeper) CreateModuleAccount(ctx sdk.Context) {
	moduleAcc := authtypes.NewEmptyModuleAccount(types.ModuleName, authtypes.Minter)

	k.ak.SetModuleAccount(ctx, moduleAcc)
}

func (k Keeper) ListUnBiddedAuctions(ctx sdk.Context) (auctions []types.Auction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var auction types.Auction
		k.cdc.MustUnmarshal(iterator.Value(), &auction)

		if strings.EqualFold(auction.BidderAddress, "") {
			auctions = append(auctions, auction)
		}
	}

	return
}

func (k Keeper) ListBids(ctx sdk.Context, billboardId, auctionId uint64) (bids []types.Bid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetBidKeyPrefix(billboardId, auctionId))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bid types.Bid
		k.cdc.MustUnmarshal(iterator.Value(), &bid)
		bids = append(bids, bid)
	}

	return
}

func (k Keeper) Bid(ctx sdk.Context, billboardId, auctionId uint64, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetBidKeyPrefix(billboardId, auctionId))

	bidderAddr := sdk.MustAccAddressFromBech32(address)

	var bid types.Bid

	bidBz := store.Get(bidderAddr) // must not null

	k.cdc.MustUnmarshal(bidBz, &bid)

	// 옥션 비더 업데이트
	auctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAuctionKeyPrefix(billboardId))
	var auction types.Auction
	auctionBz := auctionStore.Get(types.GetIDBytes(auctionId))
	k.cdc.MustUnmarshal(auctionBz, &auction)

	auction.BidderAddress = address
	auctionBz = k.cdc.MustMarshal(&auction)
	auctionStore.Set(types.GetIDBytes(auctionId), auctionBz)

	// 자금 전송
	billboardOwnerAddr := k.dk.GetOwnerAddress(ctx, billboardId)
	auctionAddr := sdk.MustAccAddressFromBech32(auction.AuctionAddress)
	k.bk.SendCoins(ctx, auctionAddr, billboardOwnerAddr, sdk.NewCoins(bid.Amount))

	// 미낙찰 자금 반환
	for _, ubid := range k.ListBids(ctx, billboardId, auctionId) {
		// 낙찰자 제외
		if !strings.EqualFold(address, ubid.SenderAddress) {
			addr := sdk.MustAccAddressFromBech32(ubid.SenderAddress)
			k.bk.SendCoins(ctx, auctionAddr, addr, sdk.NewCoins(ubid.Amount))
		}
	}

	// TODO: 비딩가격 업데이트
	// pricePerMinute 계산 필요
	// duration := auction.End - auction.Start
	// amount / duration * 60
}
