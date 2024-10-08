package client

import (
"github.com/spf13/cobra"
"github.com/cosmos/cosmos-sdk/x/gov/types"
"github.com/cosmos/cosmos-sdk/client"


)
func GetQueryCmd() *cobra.Command {
    govQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the governance module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	govQueryCmd.AddCommand(
		GetCmdQueryMinimalDeposit(),,
		
	)

	return govQueryCmd


}