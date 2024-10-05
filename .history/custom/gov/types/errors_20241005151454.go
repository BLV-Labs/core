package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrQueryExchangeRateUusdFail := sdkerrors.Register(ModuleName, 2, "unknown proposal")
)