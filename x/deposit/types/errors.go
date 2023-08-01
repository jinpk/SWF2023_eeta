package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/deposit module sentinel errors
var (
	ErrUnauthMint = sdkerrors.Register(ModuleName, 1, "unauthorized sender address for minter")
	ErrSample     = sdkerrors.Register(ModuleName, 1100, "sample error")
)
