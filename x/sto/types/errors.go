package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sto module sentinel errors
var (
	ErrFailedStoFunding         = sdkerrors.Register(ModuleName, 1, "sto funding 입금 실패")
	ErrInvalidShareTotal        = sdkerrors.Register(ModuleName, 2, "share total은 100이여야 합니다.")
	ErrInsufiendOrganazierShare = sdkerrors.Register(ModuleName, 3, "초기 펀드 잔고가 부족합니다.")
	ErrAllFunded                = sdkerrors.Register(ModuleName, 4, "fund 완료 sto")
	ErrSample                   = sdkerrors.Register(ModuleName, 1100, "sample error")
)
