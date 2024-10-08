package client_test

import (
	"fmt"
	"encoding/base64"
	"github.com/cosmos/cosmos-sdk/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/cosmos/cosmos-sdk/client/flags"
	v2lunc1cli "github.com/classic-terra/core/v3/custom/gov/client/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/classic-terra/core/v3/custom/gov/types"

)


func (s *CLITestSuite) TestNewCmdSubmitProposal() {
	val := testutil.CreateKeyringAccounts(s.T(), s.kr, 1)

	// Create a legacy proposal JSON, make sure it doesn't pass this new CLI
	// command.
	invalidProp := `{
		"title": "",
		"description": "Where is the title!?",
		"type": "Text",
		"deposit": "-324foocoin"
	}`
	invalidPropFile := testutil.WriteToNewTempFile(s.T(), invalidProp)
	defer invalidPropFile.Close()

	// Create a valid new proposal JSON.
	propMetadata := []byte{42}
	validProp := fmt.Sprintf(`
	{
		"messages": [
			{
				"@type": "/cosmos.gov.v1.MsgExecLegacyContent",
				"authority": "%s",
				"content": {
					"@type": "/cosmos.gov.v1beta1.TextProposal",
					"title": "My awesome title",
					"description": "My awesome description"
				}
			}
		],
		"title": "My awesome title",
		"summary": "My awesome description",
		"metadata": "%s",
		"deposit": "%s"
	}`, authtypes.NewModuleAddress(types.ModuleName), base64.StdEncoding.EncodeToString(propMetadata), sdk.NewCoin("stake", sdk.NewInt(5431)))
	validPropFile := testutil.WriteToNewTempFile(s.T(), validProp)
	defer validPropFile.Close()

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
		respType  proto.Message
	}{
		{
			"invalid proposal",
			[]string{
				invalidPropFile.Name(),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			true, nil,
		},
		{
			"valid proposal",
			[]string{
				validPropFile.Name(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val[0].Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))).String()),
			},
			false, &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := v2lunc1cli.NewCmdSubmitProposal()

			out, err := clitestutil.ExecTestCLICmd(s.baseCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(s.baseCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}