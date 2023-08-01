package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"eeta/x/billboard/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
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
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetOwnerAddress(ctx sdk.Context, billboardId uint64) sdk.AccAddress {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BillboardKey))
	bz := store.Get(types.GetBillboardIDBytes(billboardId))
	var billboard types.Billboard
	k.cdc.MustUnmarshal(bz, &billboard)

	return sdk.MustAccAddressFromBech32(billboard.OwnerAddress)
}

func (k Keeper) GetNextBillboardId(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BillboardKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var count uint64 = 1
	for ; iterator.Valid(); iterator.Next() {
		count++
	}

	return count
}

func (k Keeper) Has(ctx sdk.Context, billboardID uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BillboardKey))
	return store.Has(types.GetBillboardIDBytes(billboardID))
}

func (k Keeper) GetFinalBidPricePerMinute(ctx sdk.Context, billboardId uint64) sdk.Coin {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BillboardKey))
	var billboard types.Billboard
	k.cdc.MustUnmarshal(store.Get(types.GetBillboardIDBytes(billboardId)), &billboard)
	return billboard.FinalBidPricePerMinute
}
