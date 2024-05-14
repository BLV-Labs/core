package v4

import (
	"github.com/classic-terra/core/v3/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
)

const UpgradeName = "v4"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV4UpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{Added: []string{icahosttypes.StoreKey}},
}
