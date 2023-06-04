package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateUser = "create_user"

var _ sdk.Msg = &MsgCreateUser{}

func NewMsgCreateUser(creator string, name string, password string) *MsgCreateUser {
	return &MsgCreateUser{
		Name:     name,
		Password: password,
	}
}

func (msg *MsgCreateUser) Route() string {
	return RouterKey
}

func (msg *MsgCreateUser) Type() string {
	return TypeMsgCreateUser
}

func (msg *MsgCreateUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgCreateUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUser) ValidateBasic() error {
	return nil
}
