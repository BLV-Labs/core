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
		GetCmdQueryMinimalDeposit(),	
	)

	return govQueryCmd

}

func GetCmdQueryMinimalDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal [proposal-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query details of a single proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details for a proposal. You can find the
proposal-id by running "%s query gov proposals".

Example:
$ %s query gov proposal 1
`,
				version.AppName, version.AppName,
			),
		),

}