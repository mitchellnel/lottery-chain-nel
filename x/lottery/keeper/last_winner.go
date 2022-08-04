package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
)

// SetLastWinner set lastWinner in the store
func (k Keeper) SetLastWinner(ctx sdk.Context, lastWinner types.LastWinner) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastWinnerKey))
	b := k.cdc.MustMarshal(&lastWinner)
	store.Set([]byte{0}, b)
}

// GetLastWinner returns lastWinner
func (k Keeper) GetLastWinner(ctx sdk.Context) (val types.LastWinner, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastWinnerKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLastWinner removes lastWinner from the store
func (k Keeper) RemoveLastWinner(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastWinnerKey))
	store.Delete([]byte{0})
}
