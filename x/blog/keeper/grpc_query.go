package keeper

import (
	"github.com/zireael26/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
