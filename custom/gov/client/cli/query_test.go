package cli_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	rpcclientmock "github.com/cometbft/cometbft/rpc/client/mock"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutilmod "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govclitestutil "github.com/cosmos/cosmos-sdk/x/gov/client/testutil"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

type CLITestSuite struct {
	suite.Suite

	kr        keyring.Keyring
	encCfg    testutilmod.TestEncodingConfig
	baseCtx   client.Context
	clientCtx client.Context
}

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}

func (s *CLITestSuite) SetupSuite() {
	s.encCfg = testutilmod.MakeTestEncodingConfig(gov.AppModuleBasic{})
	s.kr = keyring.NewInMemory(s.encCfg.Codec)
	s.baseCtx = client.Context{}.
		WithKeyring(s.kr).
		WithTxConfig(s.encCfg.TxConfig).
		WithCodec(s.encCfg.Codec).
		WithClient(clitestutil.MockTendermintRPC{Client: rpcclientmock.Client{}}).
		WithAccountRetriever(client.MockAccountRetriever{}).
		WithOutput(io.Discard).
		WithChainID("test-chain")

	var outBuf bytes.Buffer
	ctxGen := func() client.Context {
		bz, _ := s.encCfg.Codec.Marshal(&sdk.TxResponse{})
		c := clitestutil.NewMockTendermintRPC(abci.ResponseQuery{
			Value: bz,
		})
		return s.baseCtx.WithClient(c)
	}
	s.clientCtx = ctxGen().WithOutput(&outBuf)

	val := testutil.CreateKeyringAccounts(s.T(), s.kr, 1)

	// create a proposal with deposit
	_, err := govclitestutil.MsgSubmitLegacyProposal(s.clientCtx, val[0].Address.String(),
		"Text Proposal 1", "Where is the title!?", v1beta1.ProposalTypeText,
		fmt.Sprintf("--%s=%s", cli.FlagDeposit, sdk.NewCoin("stake", v1.DefaultMinDepositTokens).String()))
	s.Require().NoError(err)

	// vote for proposal
	_, err = govclitestutil.MsgVote(s.clientCtx, val[0].Address.String(), "1", "yes")
	s.Require().NoError(err)

	// create a proposal without deposit
	_, err = govclitestutil.MsgSubmitLegacyProposal(s.clientCtx, val[0].Address.String(),
		"Text Proposal 2", "Where is the title!?", v1beta1.ProposalTypeText)
	s.Require().NoError(err)

	// create a proposal3 with deposit
	_, err = govclitestutil.MsgSubmitLegacyProposal(s.clientCtx, val[0].Address.String(),
		"Text Proposal 3", "Where is the title!?", v1beta1.ProposalTypeText,
		fmt.Sprintf("--%s=%s", cli.FlagDeposit, sdk.NewCoin("stake", v1.DefaultMinDepositTokens).String()))
	s.Require().NoError(err)

	// vote for proposal3 as val
	_, err = govclitestutil.MsgVote(s.clientCtx, val[0].Address.String(), "3", "yes=0.6,no=0.3,abstain=0.05,no_with_veto=0.05")
	s.Require().NoError(err)
}

func (s *CLITestSuite) TestCmdGetProposal() {
	testCases := []struct {
		name         string
		args         []string
		expCmdOutput string
	}{
		{
			"get proposal with json response",
			[]string{
				"1",
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			"1 --output=json",
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryProposal()
			cmd.SetArgs(tc.args)
			s.Require().Contains(fmt.Sprint(cmd), strings.TrimSpace(tc.expCmdOutput))
		})
	}
}
