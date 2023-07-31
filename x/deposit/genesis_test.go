package deposit_test

import (
	"testing"

	keepertest "eeta/testutil/keeper"
	"eeta/testutil/nullify"
	"eeta/x/deposit"
	"eeta/x/deposit/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DepositKeeper(t)
	deposit.InitGenesis(ctx, *k, genesisState)
	got := deposit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
