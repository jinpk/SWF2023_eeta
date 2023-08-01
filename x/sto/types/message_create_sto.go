package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateSto = "create_sto"

var _ sdk.Msg = &MsgCreateSto{}

func NewMsgCreateSto(creator string, billboardId uint64, organizerUrl string, organizerImageUrl string, name string, start int32, end int32, userShare int32, organizerShare int32) *MsgCreateSto {
	return &MsgCreateSto{
		Creator:           creator,
		BillboardId:       billboardId,
		OrganizerUrl:      organizerUrl,
		OrganizerImageUrl: organizerImageUrl,
		Name:              name,
		Start:             start,
		End:               end,
		UserShare:         userShare,
		OrganizerShare:    organizerShare,
	}
}

func (msg *MsgCreateSto) Route() string {
	return RouterKey
}

func (msg *MsgCreateSto) Type() string {
	return TypeMsgCreateSto
}

func (msg *MsgCreateSto) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSto) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSto) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
