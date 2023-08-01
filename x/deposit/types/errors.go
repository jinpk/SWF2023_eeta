package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/deposit module sentinel errors
var (
	ErrUnauthMint           = sdkerrors.Register(ModuleName, 1, "unauthorized sender address for minter")
	ErrInsufientBurnBalance = sdkerrors.Register(ModuleName, 2, "출금가능한 잔고가 부족합니다.")
	ErrFailedBurn           = sdkerrors.Register(ModuleName, 3, "failed burn")
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
)
