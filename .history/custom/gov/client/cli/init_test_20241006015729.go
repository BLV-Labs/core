package client_test 

import (
	"encoding/base64"
	"fmt"
	"io"
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	testutilmod "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
)

type CLITestSuite struct {
	suite.Suite

	kr        keyring.Keyring
	encCfg    testutilmod.TestEncodingConfig //This holds the encoding configuration, which is crucial for marshaling and unmarshaling data for transactions and messages.
	baseCtx   client.Context //This is a base context used for all operations in the test suite.
	clientCtx client.Context  //This is a more specialized client context that can include specific settings for individual tests.
}

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite)) 
}
//run once before any tests in the suite
func (s *CLITestSuite) SetupSuite() {
	//It initializes the necessary components, such as the codec, which is used for marshaling and unmarshaling data.
	s.encCfg = testutilmod.MakeTestEncodingConfig(gov.AppModuleBasic{})


}
