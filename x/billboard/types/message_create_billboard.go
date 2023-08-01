package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateBillboard = "create_billboard"

var _ sdk.Msg = &MsgCreateBillboard{}

func NewMsgCreateBillboard(creator string, name string, description string, url string, boardType string, finalBidPricePerMinute sdk.Coin) *MsgCreateBillboard {
	return &MsgCreateBillboard{
		Creator:                creator,
		Name:                   name,
		Description:            description,
		Url:                    url,
		BoardType:              boardType,
		FinalBidPricePerMinute: finalBidPricePerMinute,
	}
}

func (msg *MsgCreateBillboard) Route() string {
	return RouterKey
}

func (msg *MsgCreateBillboard) Type() string {
	return TypeMsgCreateBillboard
}

func (msg *MsgCreateBillboard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBillboard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBillboard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
