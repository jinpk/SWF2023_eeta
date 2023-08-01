package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

func NewStoPoolAddress(stoId uint64) sdk.AccAddress {
	key := append([]byte("sto"), sdk.Uint64ToBigEndian(stoId)...)
	return address.Module(ModuleName, key)
}
