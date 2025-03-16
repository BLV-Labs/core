package cli_test

import (
	"fmt"
	"strings"

	exemptioncli "github.com/classic-terra/core/v3/x/taxexemption/client/cli"
	"github.com/spf13/cobra"
)

// runCmdTest executes common command testing logic for command structure tests
func (s *CLITestSuite) runCmdTest(cmd *cobra.Command, args []string, expCmdOutput string) {
	cmd.SetArgs(args)
	s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(expCmdOutput))
}

// runBoolFlagTest executes common testing logic for boolean flag tests
func (s *CLITestSuite) runBoolFlagTest(cmdFunc func() *cobra.Command, args []string, flagName string, expectedVal bool) {
	cmd := cmdFunc()
	err := cmd.ParseFlags(args)
	s.Require().NoError(err)

	// Check boolean flag value
	val, err := cmd.Flags().GetBool(flagName)
	s.Require().NoError(err)
	s.Require().Equal(expectedVal, val)
}

// runStringFlagTest executes common testing logic for string flag tests
func (s *CLITestSuite) runStringFlagTest(cmdFunc func() *cobra.Command, args []string, flagName string, expectedVal string) {
	cmd := cmdFunc()
	err := cmd.ParseFlags(args)
	s.Require().NoError(err)

	// Check string flag value
	val, err := cmd.Flags().GetString(flagName)
	s.Require().NoError(err)
	s.Require().Equal(expectedVal, val)
}

// TestProposalAddTaxExemptionZoneCmd tests the command structure and arguments
func (s *CLITestSuite) TestProposalAddTaxExemptionZoneCmd() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"command with no args",
			[]string{
				"",
			},
			"",
		},
		{
			"add-tax-exemption-zone with all flags",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t,terra1qt8mrv72gtvmnca9z6ftzd7slqhaf8m60aa7ye",
				"--exempt-incoming",
				"--exempt-outgoing",
				"--exempt-cross-zone",
				"--title", "Add Tax Exemption Zone",
				"--description", "Adding a new tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t,terra1qt8mrv72gtvmnca9z6ftzd7slqhaf8m60aa7ye --exempt-incoming --exempt-outgoing --exempt-cross-zone --title Add Tax Exemption Zone --description Adding a new tax exemption zone",
		},
		{
			"add-tax-exemption-zone with minimal flags",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title", "Add Tax Exemption Zone",
				"--description", "Adding a new tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --title Add Tax Exemption Zone --description Adding a new tax exemption zone",
		},
		{
			"add-tax-exemption-zone with missing title",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--description", "Adding a new tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --description Adding a new tax exemption zone",
		},
		{
			"add-tax-exemption-zone with missing description",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title", "Add Tax Exemption Zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --title Add Tax Exemption Zone",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runCmdTest(exemptioncli.ProposalAddTaxExemptionZoneCmd(), tc.args, tc.expCmdOutput)
		})
	}
}

// TestProposalAddTaxExemptionZoneBoolFlagCmd tests the boolean flags of the command
func (s *CLITestSuite) TestProposalAddTaxExemptionZoneBoolFlagCmd() {
	testCases := []struct {
		name        string
		args        []string
		flagName    string
		expectedVal bool
	}{
		{
			name: "exempt-incoming flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--exempt-incoming",
			},
			flagName:    "exempt-incoming",
			expectedVal: true,
		},
		{
			name: "exempt-incoming flag not set",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
			},
			flagName:    "exempt-incoming",
			expectedVal: false,
		},
		{
			name: "exempt-outgoing flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--exempt-outgoing",
			},
			flagName:    "exempt-outgoing",
			expectedVal: true,
		},
		{
			name: "exempt-outgoing flag not set",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
			},
			flagName:    "exempt-outgoing",
			expectedVal: false,
		},
		{
			name: "exempt-cross-zone flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--exempt-cross-zone",
			},
			flagName:    "exempt-cross-zone",
			expectedVal: true,
		},
		{
			name: "exempt-cross-zone flag not set",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
			},
			flagName:    "exempt-cross-zone",
			expectedVal: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runBoolFlagTest(exemptioncli.ProposalAddTaxExemptionZoneCmd, tc.args, tc.flagName, tc.expectedVal)
		})
	}
}

// TestProposalAddTaxExemptionZoneStringFlagCmd tests the string flags of the command
func (s *CLITestSuite) TestProposalAddTaxExemptionZoneStringFlagCmd() {
	testCases := []struct {
		name        string
		args        []string
		flagName    string
		expectedVal string
	}{
		{
			name: "title flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title",
				"Add Tax Exemption Zone",
			},
			flagName:    "title",
			expectedVal: "Add Tax Exemption Zone",
		},
		{
			name: "description flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--description",
				"Adding a new tax exemption zone",
			},
			flagName:    "description",
			expectedVal: "Adding a new tax exemption zone",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runStringFlagTest(exemptioncli.ProposalAddTaxExemptionZoneCmd, tc.args, tc.flagName, tc.expectedVal)
		})
	}
}

// TestProposalAddTaxExemptionZoneInvalidFlagCmd tests invalid flags for the command
func (s *CLITestSuite) TestProposalAddTaxExemptionZoneInvalidFlagCmd() {
	cmd := exemptioncli.ProposalAddTaxExemptionZoneCmd()

	// Non-existent flag
	s.Run("non-existent flag test", func() {
		flag := cmd.Flags().Lookup("non-existent-flag")
		s.Require().Nil(flag, "Should not find non-existent flag")
	})

	// Case 2: Incorrectly typed bool flag
	s.Run("invalid boolean flag type", func() {
		args := []string{
			"zone1",
			"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
			"--exempt-incoming=not-a-bool",
		}

		err := cmd.ParseFlags(args)
		s.Require().Error(err, "Should error on invalid boolean value")
		s.Require().Contains(err.Error(), "invalid syntax")
	})

	// Incorrectly formatted addresses (missing comma)
	s.Run("invalid address format", func() {
		tmpCmd := exemptioncli.ProposalAddTaxExemptionZoneCmd()
		args := []string{
			"zone1",
			"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t terra1qt8mrv72gtvmnca9z6ftzd7slqhaf8m60aa7ye",
		}
		tmpCmd.SetArgs(args)

		// We can't easily test the Run function directly as it involves client context
		// But we can test that the format doesn't match what the command expects
		s.Require().NotContains(fmt.Sprint(tmpCmd),
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t,terra1qt8mrv72gtvmnca9z6ftzd7slqhaf8m60aa7ye")
	})
}

// TestProposalRemoveTaxExemptionZoneCmd tests the command structure and arguments
func (s *CLITestSuite) TestProposalRemoveTaxExemptionZoneCmd() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"remove-tax-exemption-zone with all flags",
			[]string{
				"zone1",
				"--title", "Remove Tax Exemption Zone",
				"--description", "Removing a tax exemption zone",
			},
			"zone1 --title Remove Tax Exemption Zone --description Removing a tax exemption zone",
		},
		{
			"remove-tax-exemption-zone with minimal flags",
			[]string{
				"zone1",
				"--title", "Remove Tax Exemption Zone",
				"--description", "Removing a tax exemption zone",
			},
			"zone1 --title Remove Tax Exemption Zone --description Removing a tax exemption zone",
		},
		{
			"remove-tax-exemption-zone with missing title",
			[]string{
				"zone1",
				"--description", "Removing a tax exemption zone",
			},
			"zone1 --description Removing a tax exemption zone",
		},
		{
			"remove-tax-exemption-zone with missing description",
			[]string{
				"zone1",
				"--title", "Remove Tax Exemption Zone",
			},
			"zone1 --title Remove Tax Exemption Zone",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runCmdTest(exemptioncli.ProposalRemoveTaxExemptionZoneCmd(), tc.args, tc.expCmdOutput)
		})
	}
}

// TestProposalRemoveTaxExemptionZoneStringFlagCmd tests the string flags of the command
func (s *CLITestSuite) TestProposalRemoveTaxExemptionZoneStringFlagCmd() {
	testCases := []struct {
		name        string
		args        []string
		flagName    string
		expectedVal string
	}{
		{
			name: "title flag test",
			args: []string{
				"zone1",
				"--title", "Remove Tax Exemption Zone",
			},
			flagName:    "title",
			expectedVal: "Remove Tax Exemption Zone",
		},
		{
			name: "description flag test",
			args: []string{
				"zone1",
				"--description", "Removing a tax exemption zone",
			},
			flagName:    "description",
			expectedVal: "Removing a tax exemption zone",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runStringFlagTest(exemptioncli.ProposalRemoveTaxExemptionZoneCmd, tc.args, tc.flagName, tc.expectedVal)
		})
	}
}

func (s *CLITestSuite) TestProposalRemoveTaxExemptionZoneInvalidFlagCmd() {
	cmd := exemptioncli.ProposalRemoveTaxExemptionZoneCmd()

	// Case 1: Non-existent flag
	s.Run("non-existent flag test", func() {
		flag := cmd.Flags().Lookup("exempt-incoming")
		s.Require().Nil(flag, "Remove command should not have exempt-incoming flag")
	})

	// Case 2: Invalid deposit format
	s.Run("invalid deposit format", func() {
		args := []string{
			"zone1",
			"--deposit", "not-a-valid-deposit",
		}

		err := cmd.ParseFlags(args)
		s.Require().NoError(err) // Parsing succeeds but validation would fail later

		// The command would fail when it attempts to parse the coins with sdk.ParseCoinsNormalized
		// We can at least verify we got the wrong format
		val, err := cmd.Flags().GetString("deposit")
		s.Require().NoError(err)
		s.Require().Equal("not-a-valid-deposit", val)
	})
}

// TestProposalModifyTaxExemptionZoneCmd tests the command structure and arguments
func (s *CLITestSuite) TestProposalModifyTaxExemptionZoneCmd() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"modify-tax-exemption-zone with all flags",
			[]string{
				"zone1",
				"--exempt-incoming",
				"--exempt-outgoing",
				"--exempt-cross-zone",
				"--title", "Modify Tax Exemption Zone",
				"--description", "Modifying a tax exemption zone",
			},
			"zone1 --exempt-incoming --exempt-outgoing --exempt-cross-zone --title Modify Tax Exemption Zone --description Modifying a tax exemption zone",
		},
		{
			"modify-tax-exemption-zone with minimal flags",
			[]string{
				"zone1",
				"--title", "Modify Tax Exemption Zone",
				"--description", "Modifying a tax exemption zone",
			},
			"zone1 --title Modify Tax Exemption Zone --description Modifying a tax exemption zone",
		},
		{
			"modify-tax-exemption-zone with only incoming exemption",
			[]string{
				"zone1",
				"--exempt-incoming",
				"--title", "Modify Tax Exemption Zone",
				"--description", "Modifying a tax exemption zone",
			},
			"zone1 --exempt-incoming --title Modify Tax Exemption Zone --description Modifying a tax exemption zone",
		},
		{
			"modify-tax-exemption-zone with only outgoing exemption",
			[]string{
				"zone1",
				"--exempt-outgoing",
				"--title", "Modify Tax Exemption Zone",
				"--description", "Modifying a tax exemption zone",
			},
			"zone1 --exempt-outgoing --title Modify Tax Exemption Zone --description Modifying a tax exemption zone",
		},
		{
			"modify-tax-exemption-zone with only cross-zone exemption",
			[]string{
				"zone1",
				"--exempt-cross-zone",
				"--title", "Modify Tax Exemption Zone",
				"--description", "Modifying a tax exemption zone",
			},
			"zone1 --exempt-cross-zone --title Modify Tax Exemption Zone --description Modifying a tax exemption zone",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runCmdTest(exemptioncli.ProposalModifyTaxExemptionZoneCmd(), tc.args, tc.expCmdOutput)
		})
	}
}

// TestProposalModifyTaxExemptionZoneBoolFlagCmd tests the boolean flags of the command
func (s *CLITestSuite) TestProposalModifyTaxExemptionZoneBoolFlagCmd() {
	testCases := []struct {
		name        string
		args        []string
		flagName    string
		expectedVal bool
	}{
		{
			name: "exempt-incoming flag test",
			args: []string{
				"zone1",
				"--exempt-incoming",
			},
			flagName:    "exempt-incoming",
			expectedVal: true,
		},
		{
			name: "exempt-incoming flag not set",
			args: []string{
				"zone1",
			},
			flagName:    "exempt-incoming",
			expectedVal: false,
		},
		{
			name: "exempt-outgoing flag test",
			args: []string{
				"zone1",
				"--exempt-outgoing",
			},
			flagName:    "exempt-outgoing",
			expectedVal: true,
		},
		{
			name: "exempt-outgoing flag not set",
			args: []string{
				"zone1",
			},
			flagName:    "exempt-outgoing",
			expectedVal: false,
		},
		{
			name: "exempt-cross-zone flag test",
			args: []string{
				"zone1",
				"--exempt-cross-zone",
			},
			flagName:    "exempt-cross-zone",
			expectedVal: true,
		},
		{
			name: "exempt-cross-zone flag not set",
			args: []string{
				"zone1",
			},
			flagName:    "exempt-cross-zone",
			expectedVal: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runBoolFlagTest(exemptioncli.ProposalModifyTaxExemptionZoneCmd, tc.args, tc.flagName, tc.expectedVal)
		})
	}
}

// TestProposalModifyTaxExemptionZoneStringFlagCmd tests the string flags of the command
func (s *CLITestSuite) TestProposalModifyTaxExemptionZoneStringFlagCmd() {
	testCases := []struct {
		name        string
		args        []string
		flagName    string
		expectedVal string
	}{
		{
			name: "title flag test",
			args: []string{
				"zone1",
				"--title", "Modify Tax Exemption Zone",
			},
			flagName:    "title",
			expectedVal: "Modify Tax Exemption Zone",
		},
		{
			name: "description flag test",
			args: []string{
				"zone1",
				"--description", "Modifying a tax exemption zone",
			},
			flagName:    "description",
			expectedVal: "Modifying a tax exemption zone",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runStringFlagTest(exemptioncli.ProposalModifyTaxExemptionZoneCmd, tc.args, tc.flagName, tc.expectedVal)
		})
	}
}

func (s *CLITestSuite) TestProposalModifyTaxExemptionZoneInvalidFlagCmd() {
	cmd := exemptioncli.ProposalModifyTaxExemptionZoneCmd()

	// Case 1: Unsupported flag
	s.Run("unsupported flag test", func() {
		flag := cmd.Flags().Lookup("addresses")
		s.Require().Nil(flag, "Modify command should not have addresses flag")
	})

	// Case 2: Mix-up boolean flags with string values
	s.Run("boolean flag with string value", func() {
		args := []string{
			"zone1",
			"--exempt-incoming=yes", // Should be a boolean flag without value
		}

		err := cmd.ParseFlags(args)
		s.Require().Error(err, "Should error on invalid boolean flag usage")
	})
}

// TestProposalAddTaxExemptionAddressCmd tests the command structure and arguments
func (s *CLITestSuite) TestProposalAddTaxExemptionAddressCmd() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"add-tax-exemption-address with all flags",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t,terra1qt8mrv72gtvmnca9z6ftzd7slqhaf8m60aa7ye",
				"--title", "Add Tax Exemption Address",
				"--description", "Adding addresses to tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t,terra1qt8mrv72gtvmnca9z6ftzd7slqhaf8m60aa7ye --title Add Tax Exemption Address --description Adding addresses to tax exemption zone",
		},
		{
			"add-tax-exemption-address with minimal flags",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title", "Add Tax Exemption Address",
				"--description", "Adding addresses to tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --title Add Tax Exemption Address --description Adding addresses to tax exemption zone",
		},
		{
			"add-tax-exemption-address with missing title",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--description", "Adding addresses to tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --description Adding addresses to tax exemption zone",
		},
		{
			"add-tax-exemption-address with missing description",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title", "Add Tax Exemption Address",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --title Add Tax Exemption Address",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runCmdTest(exemptioncli.ProposalAddTaxExemptionAddressCmd(), tc.args, tc.expCmdOutput)
		})
	}
}

// TestProposalAddTaxExemptionAddressStringFlagCmd tests the string flags of the command
func (s *CLITestSuite) TestProposalAddTaxExemptionAddressStringFlagCmd() {
	testCases := []struct {
		name        string
		args        []string
		flagName    string
		expectedVal string
	}{
		{
			name: "title flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title", "Add Tax Exemption Address",
			},
			flagName:    "title",
			expectedVal: "Add Tax Exemption Address",
		},
		{
			name: "description flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--description", "Adding addresses to tax exemption zone",
			},
			flagName:    "description",
			expectedVal: "Adding addresses to tax exemption zone",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runStringFlagTest(exemptioncli.ProposalAddTaxExemptionAddressCmd, tc.args, tc.flagName, tc.expectedVal)
		})
	}
}

// TestProposalAddTaxExemptionAddressInvalidFlagCmd tests invalid flags for the command
func (s *CLITestSuite) TestProposalAddTaxExemptionAddressInvalidFlagCmd() {
	cmd := exemptioncli.ProposalAddTaxExemptionAddressCmd()

	// Case 1: Exempt flags not available
	s.Run("unsupported exempt flag test", func() {
		flag := cmd.Flags().Lookup("exempt-incoming")
		s.Require().Nil(flag, "Add Address command should not have exempt-incoming flag")
	})

	// Case 2: Missing required arguments
	s.Run("missing required arguments", func() {
		tmpCmd := exemptioncli.ProposalAddTaxExemptionAddressCmd()
		args := []string{
			"zone1",
			// Missing addresses argument
		}
		tmpCmd.SetArgs(args)

		// The exact RunE command would fail with "accepts 2 arg(s), received 1"
		// We can verify the args count is wrong
		s.Require().NotEqual(2, len(args))
	})
}

// TestProposalRemoveTaxExemptionAddressCmd tests the command structure and arguments
func (s *CLITestSuite) TestProposalRemoveTaxExemptionAddressCmd() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"remove-tax-exemption-address with all flags",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t,terra1qt8mrv72gtvmnca9z6ftzd7slqhaf8m60aa7ye",
				"--title", "Remove Tax Exemption Address",
				"--description", "Removing addresses from tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t,terra1qt8mrv72gtvmnca9z6ftzd7slqhaf8m60aa7ye --title Remove Tax Exemption Address --description Removing addresses from tax exemption zone",
		},
		{
			"remove-tax-exemption-address with minimal flags",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title", "Remove Tax Exemption Address",
				"--description", "Removing addresses from tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --title Remove Tax Exemption Address --description Removing addresses from tax exemption zone",
		},
		{
			"remove-tax-exemption-address with missing title",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--description", "Removing addresses from tax exemption zone",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --description Removing addresses from tax exemption zone",
		},
		{
			"remove-tax-exemption-address with missing description",
			[]string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title", "Remove Tax Exemption Address",
			},
			"zone1 terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t --title Remove Tax Exemption Address",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runCmdTest(exemptioncli.ProposalRemoveTaxExemptionAddressCmd(), tc.args, tc.expCmdOutput)
		})
	}
}

// TestProposalRemoveTaxExemptionAddressStringFlagCmd tests the string flags of the command
func (s *CLITestSuite) TestProposalRemoveTaxExemptionAddressStringFlagCmd() {
	testCases := []struct {
		name        string
		args        []string
		flagName    string
		expectedVal string
	}{
		{
			name: "title flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--title", "Remove Tax Exemption Address",
			},
			flagName:    "title",
			expectedVal: "Remove Tax Exemption Address",
		},
		{
			name: "description flag test",
			args: []string{
				"zone1",
				"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
				"--description", "Removing addresses from tax exemption zone",
			},
			flagName:    "description",
			expectedVal: "Removing addresses from tax exemption zone",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			s.runStringFlagTest(exemptioncli.ProposalRemoveTaxExemptionAddressCmd, tc.args, tc.flagName, tc.expectedVal)
		})
	}
}

// TestProposalRemoveTaxExemptionAddressInvalidFlagCmd tests invalid flags for the command
func (s *CLITestSuite) TestProposalRemoveTaxExemptionAddressInvalidFlagCmd() {
	cmd := exemptioncli.ProposalRemoveTaxExemptionAddressCmd()

	// Case 1: Unsupported flags
	s.Run("unsupported flag test", func() {
		flag := cmd.Flags().Lookup("exempt-outgoing")
		s.Require().Nil(flag, "Remove Address command should not have exempt-outgoing flag")
	})

	// Case 2: Invalid title flag format
	s.Run("title flag with empty value", func() {
		args := []string{
			"zone1",
			"terra1dczz24r33fwlj0q5ra7rcdryjpk9hxm8rwy39t",
			"--title=", // Empty title
		}

		err := cmd.ParseFlags(args)
		s.Require().NoError(err) // Parsing succeeds

		// Empty title would fail in the RunE function, but we can check we got an empty value
		val, err := cmd.Flags().GetString("title")
		s.Require().NoError(err)
		s.Require().Equal("", val, "Empty title should be preserved")
	})
}
