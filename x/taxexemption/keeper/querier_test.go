package keeper_test

import (
	"testing"

	ultil "github.com/classic-terra/core/v3/x/taxexemption/keeper"
	"github.com/classic-terra/core/v3/x/taxexemption/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestQueryTaxable(t *testing.T) {
	input := ultil.CreateTestInput(t)
	ctx := sdk.WrapSDKContext(input.Ctx)
	querier := ultil.NewQuerier(input.TaxExemptionKeeper)
	zoneName := "zoneName"
	fAddr := "terra10d07y265gmmuvt4z0w9aw880jnsr700juxf95n"

	zone_test := types.Zone{
		Name:      zoneName,
		Outgoing:  true,
		Incoming:  true,
		CrossZone: true,
	}
	input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, zone_test)
	input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zoneName, fAddr)

	//Case 1: Empty request
	_, err := querier.Taxable(ctx, nil)
	require.Error(t, err)

	// Query to grpc
	res, err := querier.Taxable(ctx, &types.QueryTaxableRequest{
		FromAddress: fAddr,
	})
	require.NoError(t, err)
	require.Equal(t, false, res.Taxable)

}
