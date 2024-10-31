package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/classic-terra/core/v3/custom/gov/types/v2lunc1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	govsimulation "github.com/cosmos/cosmos-sdk/x/gov/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	DepositParamsMinUusdDeposit = "deposit_params_min_uusd_deposit"
)

// GenDepositParamsMinDeposit returns randomized DepositParamsMinDeposit
func GenDepositParamsMinUusdDeposit(r *rand.Rand) sdk.Coin {
	return sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simulation.RandIntBetween(r, 1, 1e3)))
}

// RandomizedGenState generates a random GenesisState for gov
func RandomizedGenState(simState *module.SimulationState) {
	startingProposalID := uint64(simState.Rand.Intn(100))

	var minDeposit sdk.Coins
	simState.AppParams.GetOrGenerate(
		simState.Cdc, govsimulation.DepositParamsMinDeposit, &minDeposit, simState.Rand,
		func(r *rand.Rand) { minDeposit = govsimulation.GenDepositParamsMinDeposit(r) },
	)

	var minUusdDeposit sdk.Coin
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsMinUusdDeposit, &minUusdDeposit, simState.Rand,
		func(r *rand.Rand) { minUusdDeposit = GenDepositParamsMinUusdDeposit(r) },
	)

	var depositPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, govsimulation.DepositParamsDepositPeriod, &depositPeriod, simState.Rand,
		func(r *rand.Rand) { depositPeriod = govsimulation.GenDepositParamsDepositPeriod(r) },
	)

	var minInitialDepositRatio sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, govsimulation.DepositMinInitialRatio, &minInitialDepositRatio, simState.Rand,
		func(r *rand.Rand) { minInitialDepositRatio = govsimulation.GenDepositMinInitialDepositRatio(r) },
	)

	var votingPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, govsimulation.VotingParamsVotingPeriod, &votingPeriod, simState.Rand,
		func(r *rand.Rand) { votingPeriod = govsimulation.GenVotingParamsVotingPeriod(r) },
	)

	var quorum sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, govsimulation.TallyParamsQuorum, &quorum, simState.Rand,
		func(r *rand.Rand) { quorum = govsimulation.GenTallyParamsQuorum(r) },
	)

	var threshold sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, govsimulation.TallyParamsThreshold, &threshold, simState.Rand,
		func(r *rand.Rand) { threshold = govsimulation.GenTallyParamsThreshold(r) },
	)

	var veto sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, govsimulation.TallyParamsVeto, &veto, simState.Rand,
		func(r *rand.Rand) { veto = govsimulation.GenTallyParamsVeto(r) },
	)

	govGenesis := v2lunc1.NewGenesisState(
		startingProposalID,
		v2lunc1.NewParams(minDeposit, depositPeriod, votingPeriod, quorum.String(), threshold.String(), veto.String(), minInitialDepositRatio.String(), simState.Rand.Intn(2) == 0, simState.Rand.Intn(2) == 0, simState.Rand.Intn(2) == 0, minUusdDeposit),
	)

	bz, err := json.MarshalIndent(&govGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated governance parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(govGenesis)
}
