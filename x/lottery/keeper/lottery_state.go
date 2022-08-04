package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
)

// SetLotteryState set lotteryState in the store
func (k Keeper) SetLotteryState(ctx sdk.Context, lotteryState types.LotteryState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryStateKey))
	b := k.cdc.MustMarshal(&lotteryState)
	store.Set([]byte{0}, b)
}

// GetLotteryState returns lotteryState
func (k Keeper) GetLotteryState(ctx sdk.Context) (val types.LotteryState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryStateKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLotteryState removes lotteryState from the store
func (k Keeper) RemoveLotteryState(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryStateKey))
	store.Delete([]byte{0})
}
