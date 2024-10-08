package client

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
    v2lunc1types "github.com/classic-terra/core/v3/custom/gov/types/v2lunc1"
	
)


func NewTxCmd(legacyPropCmds []*cobra.Command) *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Governance transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmdSubmitLegacyProp := NewCmdSubmitLegacyProposal()
	for _, propCmd := range legacyPropCmds {
		flags.AddTxFlagsToCmd(propCmd)
		cmdSubmitLegacyProp.AddCommand(propCmd)
	}

	govTxCmd.AddCommand(
		NewCmdDeposit(),
		NewCmdVote(),
		NewCmdWeightedVote(),
		NewCmdSubmitProposal(),
		NewCmdDraftProposal(),

		// Deprecated
		cmdSubmitLegacyProp,
	)

	return govTxCmd
}