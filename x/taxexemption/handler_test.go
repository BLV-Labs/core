package taxexemption_test

import (
	"testing"

	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	taxexemption "github.com/classic-terra/core/v3/x/taxexemption"
	ultil "github.com/classic-terra/core/v3/x/taxexemption/keeper"
	"github.com/classic-terra/core/v3/x/taxexemption/types"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"github.com/stretchr/testify/require"
)

// TestExemptionFilterOthersMsg tests handling of non-taxexemption messages
func TestExemptionFilterOthersMsg(t *testing.T) {
	input := ultil.CreateTestInput(t)
	h := taxexemption.NewHandler(input.TaxExemptionKeeper)

	// Test Case 1: Submit proposal message from govv1
	// Expected: Error - unrecognized taxexemption message
	emptyDepositCoins := sdk.NewCoins()
	govModuleAddr := input.TaxExemptionKeeper.GetAuthority()
	msgGovV1, _ := govv1.NewMsgSubmitProposal([]sdk.Msg{}, emptyDepositCoins, string(govModuleAddr), "metadata", "title", "summary")
	_, err := h(input.Ctx, msgGovV1)
	require.Error(t, err, "Should error on unrecognized taxexemption message from govv1")

	// Test Case 2: Submit proposal message from govv1beta1
	// Expected: Error - unrecognized taxexemption message
	testPubKey := secp256k1.GenPrivKey().PubKey()
	testAddress := sdk.AccAddress(testPubKey.Address())
	textProposal := govv1beta1.NewTextProposal("Test Proposal", "Test Description")
	msgGovV1Beta1, _ := govv1beta1.NewMsgSubmitProposal(textProposal, emptyDepositCoins, testAddress)
	_, err = h(input.Ctx, msgGovV1Beta1)
	require.Error(t, err, "Should error on unrecognized taxexemption message from govv1beta1")
}

// TestExemptionFilterMsg tests handling of taxexemption messages within proposals
func TestExemptionFilterMsg(t *testing.T) {
	input := ultil.CreateTestInput(t)
	h := taxexemption.NewHandler(input.TaxExemptionKeeper)
	emptyDepositCoins := sdk.NewCoins()
	govModuleAddr := input.TaxExemptionKeeper.GetAuthority()

	// Test Case 1: MsgAddTaxExemptionZone within proposal
	// Expected: Error - message validation
	msgAddZone := types.MsgAddTaxExemptionZone{
		Authority: string(govModuleAddr),
	}
	proposalMsgs := []sdk.Msg{&msgAddZone}
	submitProposal, _ := govv1.NewMsgSubmitProposal(proposalMsgs, emptyDepositCoins, string(govModuleAddr), "metadata", "Add Zone Proposal", "Add new tax exemption zone")
	_, err := h(input.Ctx, submitProposal)
	require.Error(t, err, "Should error on invalid MsgAddTaxExemptionZone")

	// Test Case 2: MsgRemoveTaxExemptionZone within proposal
	// Expected: Error - message validation
	msgRemoveZone := types.MsgRemoveTaxExemptionZone{
		Authority: string(govModuleAddr),
		Zone:      "test_zone",
	}
	proposalMsgs = []sdk.Msg{&msgRemoveZone}
	submitProposal, _ = govv1.NewMsgSubmitProposal(proposalMsgs, emptyDepositCoins, string(govModuleAddr), "metadata", "Remove Zone Proposal", "Remove tax exemption zone")
	_, err = h(input.Ctx, submitProposal)
	require.Error(t, err, "Should error on invalid MsgRemoveTaxExemptionZone")

	// Test Case 3: MsgModifyTaxExemptionZone within proposal
	// Expected: Error - message validation
	msgModifyZone := types.MsgModifyTaxExemptionZone{
		Authority: string(govModuleAddr),
		Zone:      "test_zone",
	}
	proposalMsgs = []sdk.Msg{&msgModifyZone}
	submitProposal, _ = govv1.NewMsgSubmitProposal(proposalMsgs, emptyDepositCoins, string(govModuleAddr), "metadata", "Modify Zone Proposal", "Modify tax exemption zone")
	_, err = h(input.Ctx, submitProposal)
	require.Error(t, err, "Should error on invalid MsgModifyTaxExemptionZone")

	// Test Case 4: MsgAddTaxExemptionAddress within proposal
	// Expected: Error - message validation
	msgAddAddr := types.MsgAddTaxExemptionAddress{
		Authority: string(govModuleAddr),
		Zone:      "test_zone",
	}
	proposalMsgs = []sdk.Msg{&msgAddAddr}
	submitProposal, _ = govv1.NewMsgSubmitProposal(proposalMsgs, emptyDepositCoins, string(govModuleAddr), "metadata", "Add Address Proposal", "Add tax exemption address")
	_, err = h(input.Ctx, submitProposal)
	require.Error(t, err, "Should error on invalid MsgAddTaxExemptionAddress")

	// Test Case 5: MsgRemoveTaxExemptionAddress within proposal
	// Expected: Error - message validation
	msgRemoveAddr := types.MsgRemoveTaxExemptionAddress{
		Authority: string(govModuleAddr),
		Zone:      "test_zone",
	}
	proposalMsgs = []sdk.Msg{&msgRemoveAddr}
	submitProposal, _ = govv1.NewMsgSubmitProposal(proposalMsgs, emptyDepositCoins, string(govModuleAddr), "metadata", "Remove Address Proposal", "Remove tax exemption address")
	_, err = h(input.Ctx, submitProposal)
	require.Error(t, err, "Should error on invalid MsgRemoveTaxExemptionAddress")
}

// TestExemptionMsg tests direct handling of taxexemption messages
func TestExemptionMsg(t *testing.T) {
	input := ultil.CreateTestInput(t)
	h := taxexemption.NewHandler(input.TaxExemptionKeeper)
	govModuleAddr := input.TaxExemptionKeeper.GetAuthority()

	// Test Case 1: Create a new tax exemption zone
	// Expected: Success - zone is created with specified parameters
	testPubKey := secp256k1.GenPrivKey().PubKey()
	testAddr := sdk.AccAddress(testPubKey.Address())

	msgAddZone := types.NewMsgAddTaxExemptionZone(
		"Create Test Zone",                     // title
		"Create a test zone for tax exemption", // description
		"test_zone",                            // zone name
		true,                                   // outgoing - allow outgoing transactions
		true,                                   // incoming - allow incoming transactions
		true,                                   // crossZone - allow cross-zone transactions
		[]string{testAddr.String()},            // initial addresses in the zone
		govModuleAddr,                          // authority address
	)
	_, err := h(input.Ctx, msgAddZone)
	require.NoError(t, err, "Should successfully create tax exemption zone")

	// Verify zone was created with correct parameters
	zone, err := input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, "test_zone")
	require.NoError(t, err, "Should find created zone")
	require.Equal(t, "test_zone", zone.Name, "Zone name should match")

	// Test Case 2: Modify existing tax exemption zone
	// Expected: Success - zone parameters are updated
	msgModifyZone := types.NewMsgModifyTaxExemptionZone(
		"Modify Test Zone",            // title
		"Modify test zone parameters", // description
		"test_zone",                   // zone name
		false,                         // outgoing - disable outgoing transactions
		true,                          // incoming - keep incoming transactions enabled
		false,                         // crossZone - disable cross-zone transactions
		govModuleAddr,                 // authority address
	)
	_, err = h(input.Ctx, msgModifyZone)
	require.NoError(t, err, "Should successfully modify tax exemption zone")

	// Test Case 3: Add address to existing zone
	// Expected: Success - address is added to zone
	msgAddAddr := types.NewMsgAddTaxExemptionAddress(
		"Add Test Address",          // title
		"Add address to test zone",  // description
		"test_zone",                 // zone name
		[]string{testAddr.String()}, // addresses to add
		govModuleAddr,               // authority address
	)
	_, err = h(input.Ctx, msgAddAddr)
	require.NoError(t, err, "Should successfully add address to zone")

	// Test Case 4: Remove address from zone
	// Expected: Success - address is removed from zone
	msgRemoveAddr := types.NewMsgRemoveTaxExemptionAddress(
		"Remove Test Address",       // title
		"Remove address from zone",  // description
		"test_zone",                 // zone name
		[]string{testAddr.String()}, // addresses to remove
		govModuleAddr,               // authority address
	)
	_, err = h(input.Ctx, msgRemoveAddr)
	require.NoError(t, err, "Should successfully remove address from zone")

	// Test Case 5: Remove tax exemption zone
	// Expected: Success - zone is removed
	msgRemoveZone := types.NewMsgRemoveTaxExemptionZone(
		"Remove Test Zone",            // title
		"Remove test zone completely", // description
		"test_zone",                   // zone name
		govModuleAddr,                 // authority address
	)
	_, err = h(input.Ctx, msgRemoveZone)
	require.NoError(t, err, "Should successfully remove tax exemption zone")

	// Verify zone was removed
	_, err = input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, "test_zone")
	require.Error(t, err, "Should not find removed zone")
}
