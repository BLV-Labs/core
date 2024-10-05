package types
import (
	"github.com/classic-terra/core/v3/x/market/types
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
	UUSDMinValueKeyPrefix      = []byte{0x05}
)  // UUSDMinValueKeyPrefix  stores the proposals raw bytes.

// ProposalKey gets a specific proposal from the store
func UUSDMinValueKey(proposalID uint64) []byte {
	return append(UUSDMinValueKeyPrefix, GetProposalIDBytes(proposalID)...)
}