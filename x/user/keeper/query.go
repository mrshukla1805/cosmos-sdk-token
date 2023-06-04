package keeper

import (
	"user/x/user/types"
)

var _ types.QueryServer = Keeper{}
