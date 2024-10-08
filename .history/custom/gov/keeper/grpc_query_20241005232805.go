package keeper

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.ProposalId == 0 {
		return nil, status.Error(codes.InvalidArgument, "proposal id can not be 0")
	}


}



