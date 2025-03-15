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

var govModuleAddr = "terra10d07y265gmmuvt4z0w9aw880jnsr700juxf95n"

func TestExemptionFilterOthersMsg(t *testing.T) {
	input := ultil.CreateTestInput(t)
	h := taxexemption.NewHandler(input.TaxExemptionKeeper)

	//case 1: normal submit proposal message from govv1
	// get error: unrecognized taxexemption
	depositCoins1 := sdk.NewCoins()
	msgv1, _ := govv1.NewMsgSubmitProposal([]sdk.Msg{}, depositCoins1, govModuleAddr, "metadata", "title", "summary")
	_, err := h(input.Ctx, msgv1)
	require.Error(t, err)

	//case 2: normal submit proposal message from govv1beta1
	// get error: unrecognized taxexemption
	pubKey := secp256k1.GenPrivKey().PubKey()
	address := sdk.AccAddress(pubKey.Address())
	prop1 := govv1beta1.NewTextProposal("prop1", "prop1")
	msg, _ := govv1beta1.NewMsgSubmitProposal(prop1, depositCoins1, address)
	_, err = h(input.Ctx, msg)
	require.Error(t, err)
}

func TestExemptionFilterMsg(t *testing.T) {
	input := ultil.CreateTestInput(t)
	h := taxexemption.NewHandler(input.TaxExemptionKeeper)
	depositCoins1 := sdk.NewCoins()
	govModuleAddr := input.TaxExemptionKeeper.GetAuthority()

	//case 1: MsgAddTaxExemptionZone msg
	// expect: No error
	msg := types.MsgAddTaxExemptionZone{
		Authority: string(govModuleAddr),
	}
	sdkMsg := []sdk.Msg{&msg}
	submitMsg, _ := govv1.NewMsgSubmitProposal(sdkMsg, depositCoins1, string(govModuleAddr), "metadata", "titlte", "sumary")
	_, err := h(input.Ctx, submitMsg)
	require.Error(t, err)

	//case 2: MsgRemoveTaxExemptionZone msg
	// expect: No error
	msg1 := types.MsgRemoveTaxExemptionZone{
		Authority: string(govModuleAddr),
		Zone:      "zonename",
	}
	sdkMsg = []sdk.Msg{&msg1}
	submitMsg, _ = govv1.NewMsgSubmitProposal(sdkMsg, depositCoins1, string(govModuleAddr), "metadata", "titlte", "sumary")
	_, err = h(input.Ctx, submitMsg)
	require.Error(t, err)

	//case 3: MsgModifyTaxExemptionZone msg
	// expect: No error
	msg3 := types.MsgModifyTaxExemptionZone{
		Authority: string(govModuleAddr),
		Zone:      "zonename",
	}
	sdkMsg = []sdk.Msg{&msg3}
	submitMsg, _ = govv1.NewMsgSubmitProposal(sdkMsg, depositCoins1, string(govModuleAddr), "metadata", "titlte", "sumary")
	_, err = h(input.Ctx, submitMsg)
	require.Error(t, err)

	//case 4: MsgModifyTaxExemptionZone msg
	// expect: No error
	msg4 := types.MsgAddTaxExemptionAddress{
		Zone:      "zonename",
		Authority: string(govModuleAddr),
	}
	sdkMsg = []sdk.Msg{&msg4}
	submitMsg, _ = govv1.NewMsgSubmitProposal(sdkMsg, depositCoins1, string(govModuleAddr), "metadata", "titlte", "sumary")
	_, err = h(input.Ctx, submitMsg)
	require.Error(t, err)

	//case 5: MsgRemoveTaxExemptionAddress msg
	// expect: No error
	msg5 := types.MsgRemoveTaxExemptionAddress{
		Authority: string(govModuleAddr),
		Zone:      "zonename",
	}
	sdkMsg = []sdk.Msg{&msg5}
	submitMsg, _ = govv1.NewMsgSubmitProposal(sdkMsg, depositCoins1, string(govModuleAddr), "metadata", "titlte", "sumary")
	_, err = h(input.Ctx, submitMsg)
	require.Error(t, err)
}
