package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"user/x/user/types"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateUserResponse{}, nil
}
