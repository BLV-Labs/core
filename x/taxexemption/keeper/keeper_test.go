package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/crypto/secp256k1"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/classic-terra/core/v3/x/taxexemption/types"
)

func TestTaxExemptionList(t *testing.T) {
	input := CreateTestInput(t)

	require.False(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, "", ""))
	require.Error(t, input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, "", ""))
	require.Error(t, input.TaxExemptionKeeper.RemoveTaxExemptionAddress(input.Ctx, "", ""))

	pubKey := secp256k1.GenPrivKey().PubKey()
	pubKey2 := secp256k1.GenPrivKey().PubKey()
	pubKey3 := secp256k1.GenPrivKey().PubKey()
	pubKey4 := secp256k1.GenPrivKey().PubKey()
	pubKey5 := secp256k1.GenPrivKey().PubKey()
	address := sdk.AccAddress(pubKey.Address())
	address2 := sdk.AccAddress(pubKey2.Address())
	address3 := sdk.AccAddress(pubKey3.Address())
	address4 := sdk.AccAddress(pubKey4.Address())
	address5 := sdk.AccAddress(pubKey5.Address())

	// add a zone
	input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, types.Zone{Name: "zone1", Outgoing: false, Incoming: false, CrossZone: false})
	input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, types.Zone{Name: "zone2", Outgoing: true, Incoming: false, CrossZone: false})
	input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, types.Zone{Name: "zone3", Outgoing: false, Incoming: true, CrossZone: true})

	// add an address to an invalid zone
	require.Error(t, input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, "zone4", address.String()))

	// add an address
	input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, "zone1", address.String())
	input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, "zone1", address2.String())
	input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, "zone2", address3.String())
	input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, "zone3", address5.String())

	require.True(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address.String(), address2.String()))
	require.False(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address.String(), address3.String()))
	require.False(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address.String(), address4.String()))

	// zone 2 allows outgoing, address 4 is not in a zone
	require.True(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address3.String(), address4.String()))

	require.False(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address3.String(), address.String()))
	require.False(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address5.String(), address.String()))

	// zone 3 allows incoming and cross zone
	require.True(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address.String(), address5.String()))

	// add it again
	input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, "zone1", address.String())
	require.True(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address.String(), address2.String()))

	// remove it
	input.TaxExemptionKeeper.RemoveTaxExemptionAddress(input.Ctx, "zone1", address.String())
	require.False(t, input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address.String(), address2.String()))
}

// TestAddTaxExemptionZone tests the AddTaxExemptionZone function
func TestAddTaxExemptionZone(t *testing.T) {
	input := CreateTestInput(t)

	// Define a test zone
	testZone := types.Zone{
		Name:      "test_zone",
		Outgoing:  true,
		Incoming:  true,
		CrossZone: false,
	}

	// Add the zone and verify no error
	err := input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, testZone)
	require.NoError(t, err, "Adding a zone should not error")

	// Retrieve the zone and verify it matches
	retrievedZone, err := input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, testZone.Name)
	require.NoError(t, err, "Getting the zone should not error")
	require.Equal(t, testZone.Name, retrievedZone.Name, "Zone name should match")
	require.Equal(t, testZone.Outgoing, retrievedZone.Outgoing, "Outgoing flag should match")
	require.Equal(t, testZone.Incoming, retrievedZone.Incoming, "Incoming flag should match")
	require.Equal(t, testZone.CrossZone, retrievedZone.CrossZone, "CrossZone flag should match")

	// Test adding another zone
	anotherZone := types.Zone{
		Name:      "another_zone",
		Outgoing:  false,
		Incoming:  true,
		CrossZone: true,
	}

	err = input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, anotherZone)
	require.NoError(t, err, "Adding another zone should not error")

	// Retrieve the second zone and verify
	retrievedZone, err = input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, anotherZone.Name)
	require.NoError(t, err, "Getting the second zone should not error")
	require.Equal(t, anotherZone.Name, retrievedZone.Name, "Zone name should match")
	require.Equal(t, anotherZone.Outgoing, retrievedZone.Outgoing, "Outgoing flag should match")
	require.Equal(t, anotherZone.Incoming, retrievedZone.Incoming, "Incoming flag should match")
	require.Equal(t, anotherZone.CrossZone, retrievedZone.CrossZone, "CrossZone flag should match")

	// Test overwriting an existing zone (should not error)
	modifiedZone := types.Zone{
		Name:      "test_zone", // Same name as first zone
		Outgoing:  false,
		Incoming:  false,
		CrossZone: true,
	}

	err = input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, modifiedZone)
	require.NoError(t, err, "Overwriting an existing zone should not error")

	// Verify the zone was overwritten
	retrievedZone, err = input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, modifiedZone.Name)
	require.NoError(t, err, "Getting the overwritten zone should not error")
	require.Equal(t, modifiedZone.Outgoing, retrievedZone.Outgoing, "Outgoing flag should be updated")
	require.Equal(t, modifiedZone.Incoming, retrievedZone.Incoming, "Incoming flag should be updated")
	require.Equal(t, modifiedZone.CrossZone, retrievedZone.CrossZone, "CrossZone flag should be updated")
}

// TestRemoveTaxExemptionZone tests the RemoveTaxExemptionZone function
func TestRemoveTaxExemptionZone(t *testing.T) {
	input := CreateTestInput(t)

	// Create test keys/addresses
	pubKey := secp256k1.GenPrivKey().PubKey()
	address := sdk.AccAddress(pubKey.Address())

	// Define a test zone
	testZone := types.Zone{
		Name:      "test_zone",
		Outgoing:  true,
		Incoming:  true,
		CrossZone: false,
	}

	// Add the zone
	err := input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, testZone)
	require.NoError(t, err, "Adding a zone should not error")

	// Add an address to the zone
	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, testZone.Name, address.String())
	require.NoError(t, err, "Adding an address to the zone should not error")

	// Verify the address is in the zone
	isExempt := input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address.String(), address.String())
	require.True(t, isExempt, "Address should be in the zone (exempt from tax to itself)")

	// Remove the zone
	err = input.TaxExemptionKeeper.RemoveTaxExemptionZone(input.Ctx, testZone.Name)
	require.NoError(t, err, "Removing an existing zone should not error")

	// Verify the zone is gone
	_, err = input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, testZone.Name)
	require.Error(t, err, "Getting a removed zone should error")

	// Verify the address is no longer in any zone
	isExempt = input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, address.String(), address.String())
	require.False(t, isExempt, "Address should no longer be in any zone")

	// Try to remove a non-existent zone
	err = input.TaxExemptionKeeper.RemoveTaxExemptionZone(input.Ctx, "nonexistent_zone")
	require.Error(t, err, "Removing a non-existent zone should error")
	require.Contains(t, err.Error(), "no such zone in exemption list", "Error should indicate zone doesn't exist")

	// Add multiple zones
	zone1 := types.Zone{
		Name:      "zone1",
		Outgoing:  true,
		Incoming:  false,
		CrossZone: false,
	}

	zone2 := types.Zone{
		Name:      "zone2",
		Outgoing:  false,
		Incoming:  true,
		CrossZone: false,
	}

	err = input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, zone1)
	require.NoError(t, err, "Adding zone1 should not error")

	err = input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, zone2)
	require.NoError(t, err, "Adding zone2 should not error")

	// Remove one zone
	err = input.TaxExemptionKeeper.RemoveTaxExemptionZone(input.Ctx, zone1.Name)
	require.NoError(t, err, "Removing zone1 should not error")

	// Verify zone1 is gone but zone2 remains
	_, err = input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, zone1.Name)
	require.Error(t, err, "Getting removed zone1 should error")

	retrievedZone, err := input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, zone2.Name)
	require.NoError(t, err, "Getting zone2 should not error")
	require.Equal(t, zone2.Name, retrievedZone.Name, "Zone2 should still exist")
}

// TestModifyTaxExemptionZone tests the ModifyTaxExemptionZone function
func TestModifyTaxExemptionZone(t *testing.T) {
	input := CreateTestInput(t)

	// Define a test zone
	originalZone := types.Zone{
		Name:      "test_zone",
		Outgoing:  true,
		Incoming:  false,
		CrossZone: false,
	}

	// Add the zone
	err := input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, originalZone)
	require.NoError(t, err, "Adding a zone should not error")

	// Verify the zone exists with original values
	retrievedZone, err := input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, originalZone.Name)
	require.NoError(t, err, "Getting the zone should not error")
	require.Equal(t, originalZone.Outgoing, retrievedZone.Outgoing, "Outgoing flag should match")
	require.Equal(t, originalZone.Incoming, retrievedZone.Incoming, "Incoming flag should match")
	require.Equal(t, originalZone.CrossZone, retrievedZone.CrossZone, "CrossZone flag should match")

	// Modify the zone
	modifiedZone := types.Zone{
		Name:      "test_zone", // Same name as original
		Outgoing:  false,       // Changed
		Incoming:  true,        // Changed
		CrossZone: true,        // Changed
	}

	err = input.TaxExemptionKeeper.ModifyTaxExemptionZone(input.Ctx, modifiedZone)
	require.NoError(t, err, "Modifying an existing zone should not error")

	// Verify the zone was modified
	retrievedZone, err = input.TaxExemptionKeeper.GetTaxExemptionZone(input.Ctx, modifiedZone.Name)
	require.NoError(t, err, "Getting the modified zone should not error")
	require.Equal(t, modifiedZone.Outgoing, retrievedZone.Outgoing, "Outgoing flag should be updated")
	require.Equal(t, modifiedZone.Incoming, retrievedZone.Incoming, "Incoming flag should be updated")
	require.Equal(t, modifiedZone.CrossZone, retrievedZone.CrossZone, "CrossZone flag should be updated")

	// Try to modify a non-existent zone
	nonExistentZone := types.Zone{
		Name:      "nonexistent_zone",
		Outgoing:  true,
		Incoming:  true,
		CrossZone: true,
	}

	err = input.TaxExemptionKeeper.ModifyTaxExemptionZone(input.Ctx, nonExistentZone)
	require.Error(t, err, "Modifying a non-existent zone should error")
	require.Contains(t, err.Error(), "no such zone in exemption list", "Error should indicate zone doesn't exist")
}

// TestAddTaxExemptionAddress tests the AddTaxExemptionAddress function
func TestAddTaxExemptionAddress(t *testing.T) {
	input := CreateTestInput(t)

	// Create test keys/addresses
	pubKey1 := secp256k1.GenPrivKey().PubKey()
	pubKey2 := secp256k1.GenPrivKey().PubKey()
	pubKey3 := secp256k1.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pubKey1.Address())
	addr2 := sdk.AccAddress(pubKey2.Address())
	addr3 := sdk.AccAddress(pubKey3.Address())

	// Define test zones
	zone1 := types.Zone{
		Name:      "zone1",
		Outgoing:  true,
		Incoming:  true,
		CrossZone: false,
	}

	zone2 := types.Zone{
		Name:      "zone2",
		Outgoing:  false,
		Incoming:  true,
		CrossZone: true,
	}

	// Try to add address to non-existent zone
	err := input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, "nonexistent_zone", addr1.String())
	require.Error(t, err, "Adding address to non-existent zone should error")
	require.Contains(t, err.Error(), "no such zone in exemption list", "Error should indicate zone doesn't exist")

	// Add the zones
	err = input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, zone1)
	require.NoError(t, err, "Adding zone1 should not error")

	err = input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, zone2)
	require.NoError(t, err, "Adding zone2 should not error")

	// Test adding address with invalid format
	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zone1.Name, "invalid-address")
	require.Error(t, err, "Adding invalid address should error")

	// Add address to zone1
	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zone1.Name, addr1.String())
	require.NoError(t, err, "Adding address to zone1 should not error")

	// Verify address is in zone1
	isExempt := input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr1.String(), addr1.String())
	require.True(t, isExempt, "Address1 should be exempt from tax to itself")

	// Add multiple addresses to zone2
	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zone2.Name, addr2.String())
	require.NoError(t, err, "Adding address2 to zone2 should not error")

	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zone2.Name, addr3.String())
	require.NoError(t, err, "Adding address3 to zone2 should not error")

	// Verify addresses are in zone2
	isExempt = input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr2.String(), addr2.String())
	require.True(t, isExempt, "Address2 should be exempt from tax to itself")

	isExempt = input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr3.String(), addr3.String())
	require.True(t, isExempt, "Address3 should be exempt from tax to itself")

	// Test tax exemption between addresses in the same zone
	isExempt = input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr2.String(), addr3.String())
	require.True(t, isExempt, "Addresses in the same zone should be exempt from tax to each other")

	// Test tax exemption between addresses in different zones
	// zone1: outgoing=true, incoming=true, crossZone=false
	// zone2: outgoing=false, incoming=true, crossZone=true
	// addr1 -> addr2: Exempt (zone2 has incoming and crossZone)
	isExempt = input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr1.String(), addr2.String())
	require.True(t, isExempt, "Address1 -> Address2 should be exempt (zone2 has incoming and crossZone)")

	// addr2 -> addr1: Not exempt (zone2 has crossZone but not outgoing)
	isExempt = input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr2.String(), addr1.String())
	require.False(t, isExempt, "Address2 -> Address1 should not be exempt (zone2 has crossZone but not outgoing)")

	// Add same address again (should not error)
	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zone1.Name, addr1.String())
	require.NoError(t, err, "Adding the same address again should not error")

	// Try to add address to a different zone
	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zone2.Name, addr1.String())
	require.Error(t, err, "Adding address already in zone1 to zone2 should error")
	require.Contains(t, err.Error(), "already associated with a different zone", "Error should indicate address is already in a zone")
}

// TestRemoveTaxExemptionAddress tests the RemoveTaxExemptionAddress function
func TestRemoveTaxExemptionAddress(t *testing.T) {
	input := CreateTestInput(t)

	// Create test keys/addresses
	pubKey1 := secp256k1.GenPrivKey().PubKey()
	pubKey2 := secp256k1.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pubKey1.Address())
	addr2 := sdk.AccAddress(pubKey2.Address())

	// Define test zones
	zone1 := types.Zone{
		Name:      "zone1",
		Outgoing:  true,
		Incoming:  false,
		CrossZone: false,
	}

	zone2 := types.Zone{
		Name:      "zone2",
		Outgoing:  false,
		Incoming:  true,
		CrossZone: false,
	}

	// Add the zones
	err := input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, zone1)
	require.NoError(t, err, "Adding zone1 should not error")

	err = input.TaxExemptionKeeper.AddTaxExemptionZone(input.Ctx, zone2)
	require.NoError(t, err, "Adding zone2 should not error")

	// Add addresses to the zones
	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zone1.Name, addr1.String())
	require.NoError(t, err, "Adding address to zone1 should not error")

	err = input.TaxExemptionKeeper.AddTaxExemptionAddress(input.Ctx, zone2.Name, addr2.String())
	require.NoError(t, err, "Adding address to zone2 should not error")

	// Verify addresses are in their zones
	isExempt := input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr1.String(), addr1.String())
	require.True(t, isExempt, "Address1 should be exempt from tax to itself")

	// Test removing address from zone
	err = input.TaxExemptionKeeper.RemoveTaxExemptionAddress(input.Ctx, zone1.Name, addr1.String())
	require.NoError(t, err, "Removing address from zone should not error")

	// Verify address was removed
	isExempt = input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr1.String(), addr1.String())
	require.False(t, isExempt, "Address1 should no longer be exempt after removal")

	// Second address should still be in its zone
	isExempt = input.TaxExemptionKeeper.IsExemptedFromTax(input.Ctx, addr2.String(), addr2.String())
	require.True(t, isExempt, "Address2 should still be exempt")

	// Test removing an address from the wrong zone
	err = input.TaxExemptionKeeper.RemoveTaxExemptionAddress(input.Ctx, zone1.Name, addr2.String())
	require.Error(t, err, "Removing address from wrong zone should error")

	// Test removing an address that doesn't exist in any zone
	err = input.TaxExemptionKeeper.RemoveTaxExemptionAddress(input.Ctx, zone1.Name, addr1.String())
	require.Error(t, err, "Removing an address not in any zone should error")

	// Test removing address from non-existent zone
	err = input.TaxExemptionKeeper.RemoveTaxExemptionAddress(input.Ctx, "nonexistent_zone", addr2.String())
	require.Error(t, err, "Removing address from non-existent zone should error")

	// Test removing invalid address format
	err = input.TaxExemptionKeeper.RemoveTaxExemptionAddress(input.Ctx, zone2.Name, "invalid-address")
	require.Error(t, err, "Removing invalid address should error")
}
