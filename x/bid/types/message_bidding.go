package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBidding = "bidding"

var _ sdk.Msg = &MsgBidding{}

func NewMsgBidding(creator string, billboardId uint64, auctionId uint64, amount sdk.Coin, adUrl string) *MsgBidding {
	return &MsgBidding{
		Creator:     creator,
		BillboardId: billboardId,
		AuctionId:   auctionId,
		Amount:      amount,
		AdUrl:       adUrl,
	}
}

func (msg *MsgBidding) Route() string {
	return RouterKey
}

func (msg *MsgBidding) Type() string {
	return TypeMsgBidding
}

func (msg *MsgBidding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBidding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBidding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
