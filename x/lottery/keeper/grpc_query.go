package keeper

import (
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
)

var _ types.QueryServer = Keeper{}
