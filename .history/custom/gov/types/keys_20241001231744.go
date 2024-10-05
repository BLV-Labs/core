package types
import (
	"encoding/binary"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/kv"
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

func GetProposalIDBytes(proposalID uint64) (proposalIDBz []byte) {
	proposalIDBz = make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBz, proposalID)
	return
}