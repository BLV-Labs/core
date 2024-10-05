package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	v2lunc1types "github.com/classic-terra/core/v3/custom/gov/types/v2lunc1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

type msgServer struct {
	*Keeper
	v1MsgServer govv1.MsgServer
}

// NewMsgServerImpl returns an implementation of the gov MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) v2lunc1types.MsgServer {
	return &msgServer{
		Keeper:      keeper,
		v1MsgServer: govkeeper.NewMsgServerImpl(keeper.Keeper),
	}
}

var _ v2lunc1types.MsgServer = msgServer{}

// SubmitProposal implements the MsgServer.SubmitProposal method.
func (k msgServer) SubmitProposal(goCtx context.Context, msg *govv1.MsgSubmitProposal) (*govv1.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// initialDeposit := msg.GetInitialDeposit()

	// if err := k.validateInitialDeposit(ctx, initialDeposit); err != nil {
	// 	return nil, err
	// }

	proposalMsgs, err := sdktx.GetMsgs(msg.Messages, "sdk.MsgProposal")
	if err != nil {
		return nil, err
	}

	proposer, err := sdk.AccAddressFromBech32(msg.GetProposer())
	if err != nil {
		return nil, err
	}

	proposal, err := k.Keeper.SubmitProposal(ctx, proposalMsgs, msg.Metadata, msg.Title, msg.Summary, proposer)
	if err != nil {
		return nil, err
	}

	bytes, err := proposal.Marshal()
	if err != nil {
		return nil, err
	}

	// ref: https://github.com/cosmos/cosmos-sdk/issues/9683
	ctx.GasMeter().ConsumeGas(
		3*ctx.KVGasConfig().WriteCostPerByte*uint64(len(bytes)),
		"submit proposal",
	)

	votingStarted, err := k.Keeper.AddDeposit(ctx, proposal.Id, proposer, msg.GetInitialDeposit())
	if err != nil {
		return nil, err
	}

	if votingStarted {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(govtypes.EventTypeSubmitProposal,
				sdk.NewAttribute(govtypes.AttributeKeyVotingPeriodStart, fmt.Sprintf("%d", proposal.Id)),
			),
		)
	}

	return &govv1.MsgSubmitProposalResponse{
		ProposalId: proposal.Id,
	}, nil
}

// Deposit implements the MsgServer.Deposit method.
func (k msgServer) Deposit(goCtx context.Context, msg *govv1.MsgDeposit) (*govv1.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	accAddr, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}
	votingStarted, err := k.Keeper.AddDeposit(ctx, msg.ProposalId, accAddr, msg.Amount)
	if err != nil {
		return nil, err
	}

	if votingStarted {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				govtypes.EventTypeProposalDeposit,
				sdk.NewAttribute(govtypes.AttributeKeyVotingPeriodStart, fmt.Sprintf("%d", msg.ProposalId)),
			),
		)
	}

	return &govv1.MsgDepositResponse{}, nil
}

func (k msgServer) ExecLegacyContent(goCtx context.Context, msg *govv1.MsgExecLegacyContent) (*govv1.MsgExecLegacyContentResponse, error) {
	return k.v1MsgServer.ExecLegacyContent(goCtx, msg)
}

func (k msgServer) Vote(goCtx context.Context, msg *govv1.MsgVote) (*govv1.MsgVoteResponse, error) {
	return k.v1MsgServer.Vote(goCtx, msg)
}

func (k msgServer) VoteWeighted(goCtx context.Context, msg *govv1.MsgVoteWeighted) (*govv1.MsgVoteWeightedResponse, error) {
	return k.v1MsgServer.VoteWeighted(goCtx, msg)
}

func (k msgServer) UpdateParams(goCtx context.Context, msg *v2lunc1types.MsgUpdateParams) (*v2lunc1types.MsgUpdateParamsResponse, error) {
	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &v2lunc1types.MsgUpdateParamsResponse{}, nil
}

// TODO: v1beta1 ???
