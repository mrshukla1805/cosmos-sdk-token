package user_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "user/testutil/keeper"
	"user/testutil/nullify"
	"user/x/user"
	"user/x/user/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UserKeeper(t)
	user.InitGenesis(ctx, *k, genesisState)
	got := user.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
