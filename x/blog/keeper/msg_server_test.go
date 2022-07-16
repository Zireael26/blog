package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/zireael26/blog/testutil/keeper"
	"github.com/zireael26/blog/x/blog/keeper"
	"github.com/zireael26/blog/x/blog/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
