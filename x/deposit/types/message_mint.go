package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeposit = "mint"

var _ sdk.Msg = &MsgMint{}

func NewMsgMint(sender, receipientAddress string, coin sdk.Coin) *MsgMint {
	return &MsgMint{
		Sender:            sender,
		ReceipientAddress: receipientAddress,
		Coin:              coin,
	}
}

func (msg *MsgMint) Route() string {
	return RouterKey
}

func (msg *MsgMint) Type() string {
	return TypeMsgDeposit
}

func (msg *MsgMint) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender creator address (%s)", err)
	}

	if msg.Coin.Amount.Equal(math.NewInt(0)) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "amount is zero", err)
	}
	return nil
}
