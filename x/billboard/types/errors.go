package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/billboard module sentinel errors
var (
	ErrNotFoundBillboard = sdkerrors.Register(ModuleName, 1, "not found billboard")
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
)
