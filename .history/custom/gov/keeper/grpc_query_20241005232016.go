package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	v2lunc1types "github.com/classic-terra/core/v3/custom/gov/types/v2lunc1"
	
)
var (
	_ v2lunc1types.QueryServer = queryServer{}
)
type queryServer struct{ k *Keeper }

func NewQueryServer(k *Keeper) v2lunc1types.QueryServer {
	return queryServer{k: k}
}

func (q queryServer) ProposalMinimalLUNCByUusd(ctx context.Context, req *v2lunc1types.QueryProposalRequest) (*v2lunc1types.QueryMinimalDepositProposalResponse, error) {
	// Fetch the proposal using the proposal ID
	proposal, found := q.k.GetProposal(ctx, req.proposal_id)
	proposal, found := q.k
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrUnknownProposal, "proposal with ID %d not found", req.ProposalId)
	}

	// Assuming you have logic to calculate the minimal LUNC deposit
	minDepositLUNC, err := q.k.GetMinimalDepositLUNC(ctx, req.ProposalId)
	if err != nil {
		return nil, err
	}

	// Return the minimal LUNC deposit in the response
	return &v2lunc1types.QueryProposalResponse{
		MinimalDeposit: minDepositLUNC,
	}, nil
}






