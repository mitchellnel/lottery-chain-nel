package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lottery-chain-nel/x/lottery/types"
)

// SetPlayer set a specific player in the store from its index
func (k Keeper) SetPlayer(ctx sdk.Context, player types.Player) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKeyPrefix))
	b := k.cdc.MustMarshal(&player)
	store.Set(types.PlayerKey(
		player.Address,
	), b)
}

// GetPlayer returns a player from its index
func (k Keeper) GetPlayer(
	ctx sdk.Context,
	address string,

) (val types.Player, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKeyPrefix))

	b := store.Get(types.PlayerKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePlayer removes a player from the store
func (k Keeper) RemovePlayer(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKeyPrefix))
	store.Delete(types.PlayerKey(
		address,
	))
}

// GetAllPlayer returns all player
func (k Keeper) GetAllPlayer(ctx sdk.Context) (list []types.Player) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Player
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
