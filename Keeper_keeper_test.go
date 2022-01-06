package keeper_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/dispensation/types"
	"github.com/cosmos/cosmos-sdk/codec"
	
)




func TestKeeper_Logger(t *testing.T){
    key := types.StoreKey
	cdc := codec.BinaryCodec(types.ModuleCdc)
    bankKeeper:= types.AccountKeeper()


	
}
