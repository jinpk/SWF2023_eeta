package keeper

import (
	"eeta/x/deposit/types"
)

var _ types.QueryServer = Keeper{}
