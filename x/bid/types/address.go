package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

func NewAuctionPoolAddress(auctionId uint64) sdk.AccAddress {
	key := append([]byte("auction"), sdk.Uint64ToBigEndian(auctionId)...)
	return address.Module(ModuleName, key)
}
