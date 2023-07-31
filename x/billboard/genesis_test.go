package billboard_test

import (
	"testing"

	keepertest "eeta/testutil/keeper"
	"eeta/testutil/nullify"
	"eeta/x/billboard"
	"eeta/x/billboard/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BillboardKeeper(t)
	billboard.InitGenesis(ctx, *k, genesisState)
	got := billboard.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
