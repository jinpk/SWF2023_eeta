package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bid module sentinel errors
var (
	ErrExistAuctionTime = sdkerrors.Register(ModuleName, 1, "해당 빌보드의 광고 시간은 이미 선점되었습니다.")
	ErrInsufientBalance = sdkerrors.Register(ModuleName, 2, "낙찰 가격의 잔고가 없습니다")
	ErrSample           = sdkerrors.Register(ModuleName, 1100, "sample error")
)
