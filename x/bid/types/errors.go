package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bid module sentinel errors
var (
	ErrExistAuctionTime = sdkerrors.Register(ModuleName, 1, "해당 빌보드의 광고 시간은 이미 선점되었습니다.")
	ErrInsufientBalance = sdkerrors.Register(ModuleName, 2, "낙찰 가격의 잔고가 없습니다")
	ErrNotFoundAuction  = sdkerrors.Register(ModuleName, 3, "옥션이 없습니다")
	ErrAuctionBidded    = sdkerrors.Register(ModuleName, 3, "낙찰된 경매입니다")
	ErrAlreadyBidded    = sdkerrors.Register(ModuleName, 3, "이미 입찰 하였습니다")
	ErrSample           = sdkerrors.Register(ModuleName, 1100, "sample error")
)
