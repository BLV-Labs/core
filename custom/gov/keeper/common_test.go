package keeper_test

import (
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/math"
	"cosmossdk.io/store"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/golang/mock/gomock"

	"github.com/classic-terra/core/v3/custom/gov/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	v2lunc1types "github.com/classic-terra/core/v3/custom/gov/types/v2lunc1"
	oraclekeeper "github.com/classic-terra/core/v3/x/oracle/keeper"
	oracletypes "github.com/classic-terra/core/v3/x/oracle/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	_, _, addr   = testdata.KeyTestPubAddr()
	govAcct      = authtypes.NewModuleAddress(types.ModuleName)
	TestProposal = getTestProposal()
)

// getTestProposal creates and returns a test proposal message.
func getTestProposal() []sdk.Msg {
	legacyProposalMsg, err := v1.NewLegacyContent(v1beta1.NewTextProposal("Title", "description"), authtypes.NewModuleAddress(types.ModuleName).String())
	if err != nil {
		panic(err)
	}

	return []sdk.Msg{
		banktypes.NewMsgSend(govAcct, addr, sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000)))),
		legacyProposalMsg,
	}
}

// setupGovKeeper creates a govKeeper as well as all its dependencies.
func setupGovKeeper(t *testing.T) (
	*keeper.Keeper,
	*govtestutil.MockAccountKeeper,
	*govtestutil.MockBankKeeper,
	*govtestutil.MockStakingKeeper,
	moduletestutil.TestEncodingConfig,
	sdk.Context,
) {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(t, key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()
	v1.RegisterInterfaces(encCfg.InterfaceRegistry)
	v1beta1.RegisterInterfaces(encCfg.InterfaceRegistry)
	banktypes.RegisterInterfaces(encCfg.InterfaceRegistry)

	// Create MsgServiceRouter, but don't populate it before creating the gov
	// keeper.
	msr := baseapp.NewMsgServiceRouter()

	// gomock initializations
	ctrl := gomock.NewController(t)
	acctKeeper := govtestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := govtestutil.NewMockBankKeeper(ctrl)
	stakingKeeper := govtestutil.NewMockStakingKeeper(ctrl)
	acctKeeper.EXPECT().GetModuleAddress(types.ModuleName).Return(govAcct).AnyTimes()
	acctKeeper.EXPECT().GetModuleAccount(gomock.Any(), types.ModuleName).Return(authtypes.NewEmptyModuleAccount(types.ModuleName)).AnyTimes()
	trackMockBalances(bankKeeper)
	stakingKeeper.EXPECT().TokensFromConsensusPower(ctx, gomock.Any()).DoAndReturn(func(ctx sdk.Context, power int64) math.Int {
		return sdk.TokensFromConsensusPower(power, math.NewIntFromUint64(1000000))
	}).AnyTimes()
	stakingKeeper.EXPECT().BondDenom(ctx).Return("stake").AnyTimes()
	stakingKeeper.EXPECT().IterateBondedValidatorsByPower(gomock.Any(), gomock.Any()).AnyTimes()
	stakingKeeper.EXPECT().IterateDelegations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	stakingKeeper.EXPECT().TotalBondedTokens(gomock.Any()).Return(math.NewInt(10000000)).AnyTimes()

	oracleKeeper := InitOracleKeeper(acctKeeper, bankKeeper, distrKeeper, stakingKeeper)
	// Gov keeper initializations
	govKeeper := keeper.NewKeeper(encCfg.Codec, key, acctKeeper, bankKeeper, stakingKeeper, oracleKeeper, msr, types.DefaultConfig(), govAcct.String())
	govKeeper.SetProposalID(ctx, 1)
	govRouter := v1beta1.NewRouter() // Also register legacy gov handlers to test them too.
	govRouter.AddRoute(types.RouterKey, v1beta1.ProposalHandler)
	govKeeper.SetLegacyRouter(govRouter)
	govKeeper.SetParams(ctx, v2lunc1types.DefaultParams())

	// Register all handlers for the MegServiceRouter.
	msr.SetInterfaceRegistry(encCfg.InterfaceRegistry)
	v2lunc1types.RegisterMsgServer(msr, keeper.NewMsgServerImpl(govKeeper))
	banktypes.RegisterMsgServer(msr, nil) // Nil is fine here as long as we never execute the proposal's Msgs.

	return govKeeper, acctKeeper, bankKeeper, stakingKeeper, oracleKeeper, encCfg, ctx
}

func InitOracleKeeper(accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	distrKeeper types.DistributionKeeper,
	stakingKeeper types.StakingKeeper) {
	encodingConfig := moduletestutil.MakeTestEncodingConfig()
	keyOracle := sdk.NewKVStoreKey(types.StoreKey)
	keyParams := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tKeyParams := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ctx := sdk.NewContext(ms, tmproto.Header{Time: time.Now().UTC()}, false, log.NewNopLogger())

	// Gov keeper initializations
	appCodec, legacyAmino := encodingConfig.Codec, encodingConfig.Amino
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, keyParams, tKeyParams)

	oracleKeeper := oraclekeeper.NewKeeper(
		appCodec,
		keyOracle,
		paramsKeeper.Subspace(oracletypes.ModuleName),
		accountKeeper,
		bankKeeper,
		distrKeeper,
		stakingKeeper,
		distrtypes.ModuleName,
	)
	oracleDefaultParams := oracletypes.DefaultParams()
	oracleKeeper.SetParams(ctx, oracleDefaultParams)

	for _, denom := range oracleDefaultParams.Whitelist {
		oracleKeeper.SetTobinTax(ctx, denom.Name, denom.TobinTax)
	}
	// End Gov keeper initializations
}

// trackMockBalances sets up expected calls on the Mock BankKeeper, and also
// locally tracks accounts balances (not modules balances).
func trackMockBalances(bankKeeper *govtestutil.MockBankKeeper) {
	balances := make(map[string]sdk.Coins)

	// We don't track module account balances.
	bankKeeper.EXPECT().MintCoins(gomock.Any(), minttypes.ModuleName, gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, gomock.Any()).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToModule(gomock.Any(), minttypes.ModuleName, types.ModuleName, gomock.Any()).AnyTimes()

	// But we do track normal account balances.
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), types.ModuleName, gomock.Any()).DoAndReturn(func(_ sdk.Context, sender sdk.AccAddress, _ string, coins sdk.Coins) error {
		newBalance, negative := balances[sender.String()].SafeSub(coins...)
		if negative {
			return fmt.Errorf("not enough balance")
		}
		balances[sender.String()] = newBalance
		return nil
	}).AnyTimes()
	bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ sdk.Context, module string, rcpt sdk.AccAddress, coins sdk.Coins) error {
		balances[rcpt.String()] = balances[rcpt.String()].Add(coins...)
		return nil
	}).AnyTimes()
	bankKeeper.EXPECT().GetAllBalances(gomock.Any(), gomock.Any()).DoAndReturn(func(_ sdk.Context, addr sdk.AccAddress) sdk.Coins {
		return balances[addr.String()]
	}).AnyTimes()
}
