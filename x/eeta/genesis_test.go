package eeta_test

import (
	"testing"

	keepertest "eeta/testutil/keeper"
	"eeta/testutil/nullify"
	"eeta/x/eeta"
	"eeta/x/eeta/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EetaKeeper(t)
	eeta.InitGenesis(ctx, *k, genesisState)
	got := eeta.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
