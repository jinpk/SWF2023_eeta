package keeper_test

import (
	"testing"

	testkeeper "eeta/testutil/keeper"
	"eeta/x/bid/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BidKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.WaitBlock, k.WaitBlock(ctx))
}
