package types

import (
	"cosmossdk.io/collections"
)

const (
	// ModuleName is the name of the module
	ModuleName = "gov"

	// StoreKey is the store key string for gov
	StoreKey = ModuleName

	// RouterKey is the message route for gov
	RouterKey = ModuleName
)

var (
	UUSDMinValueKey           = collections.NewPrefix(0)  // ProposalsKeyPrefix stores the proposals raw bytes.
)