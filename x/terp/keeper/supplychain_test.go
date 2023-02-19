package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/terpnetwork/terp-core/testutil/keeper"
	"github.com/terpnetwork/terp-core/testutil/nullify"
	"github.com/terpnetwork/terp-core/x/terp/keeper"
	"github.com/terpnetwork/terp-core/x/terp/types"
)

func createNSupplychain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Supplychain {
	items := make([]types.Supplychain, n)
	for i := range items {
		items[i].Id = keeper.AppendSupplychain(ctx, items[i])
	}
	return items
}

func TestSupplychainGet(t *testing.T) {
	keeper, ctx := keepertest.TerpKeeper(t)
	items := createNSupplychain(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSupplychain(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSupplychainRemove(t *testing.T) {
	keeper, ctx := keepertest.TerpKeeper(t)
	items := createNSupplychain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSupplychain(ctx, item.Id)
		_, found := keeper.GetSupplychain(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSupplychainGetAll(t *testing.T) {
	keeper, ctx := keepertest.TerpKeeper(t)
	items := createNSupplychain(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSupplychain(ctx)),
	)
}

func TestSupplychainCount(t *testing.T) {
	keeper, ctx := keepertest.TerpKeeper(t)
	items := createNSupplychain(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSupplychainCount(ctx))
}