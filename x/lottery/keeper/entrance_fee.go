package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
)

// SetEntranceFee set entranceFee in the store
func (k Keeper) SetEntranceFee(ctx sdk.Context, entranceFee types.EntranceFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntranceFeeKey))
	b := k.cdc.MustMarshal(&entranceFee)
	store.Set([]byte{0}, b)
}

// GetEntranceFee returns entranceFee
func (k Keeper) GetEntranceFee(ctx sdk.Context) (val types.EntranceFee, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntranceFeeKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEntranceFee removes entranceFee from the store
func (k Keeper) RemoveEntranceFee(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EntranceFeeKey))
	store.Delete([]byte{0})
}
