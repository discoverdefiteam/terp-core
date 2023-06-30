package v2

import (
	"fmt"

	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v7/router/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/terpnetwork/terp-core/v2/app"
	"github.com/terpnetwork/terp-core/v2/app/upgrades"
)

func CreateV2UpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
	keepers *app.TerpApp,
) upgradetypes.UpgradeHandler {

	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// transfer module consensus version has been bumped to 2
		// the above is https://github.com/cosmos/ibc-go/blob/v5.1.0/docs/migrations/v3-to-v4.md
		logger := ctx.Logger().With("upgrade", UpgradeName)

		nativeDenom := upgrades.GetChainsDenomToken(ctx.ChainID())
		logger.Info(fmt.Sprintf("With native denom %s", nativeDenom))

		// Run migrations
		versionMap, err := mm.RunMigrations(ctx, cfg, vm)

		// New modules run AFTER the migrations, so to set the correct params after the default.

		// Packet Forward middleware initial params
		keepers.PacketForwardKeeper.SetParams(ctx, packetforwardtypes.DefaultParams())

		return versionMap, err
	}

}