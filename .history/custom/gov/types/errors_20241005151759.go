package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	ErrQueryExchangeRateUusdFail = sdkerrors.Register(govtypes.ModuleName, 17, "unknown proposal")
)