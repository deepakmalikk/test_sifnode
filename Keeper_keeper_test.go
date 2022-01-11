package keeper_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/dispensation/keeper"
	"github.com/Sifchain/sifnode/x/dispensation/test"
	"github.com/Sifchain/sifnode/x/dispensation/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func TestKeeper_Logger(t *testing.T) {
	app, ctx := test.CreateTestApp(false)

	key := sdk.NewKVStoreKey("rowan")
	cdc := codec.BinaryCodec(app.AppCodec())
	accountKeeper := types.AccountKeeper(app.AccountKeeper)
	bankkeeper := types.BankKeeper(app.BankKeeper)
	ps := paramtypes.Subspace{}
	result := keeper.NewKeeper(cdc, key, bankkeeper, accountKeeper, ps)

	res := result.Logger(ctx)
	t.Log(res)

}

func TestKeeper_codec(t *testing.T) {
	app, ctx := test.CreateTestApp(false)
	t.Log(ctx)
	key := sdk.NewKVStoreKey("rowan")
	cdc := codec.BinaryCodec(app.AppCodec())
	accountKeeper := types.AccountKeeper(app.AccountKeeper)
	bankkeeper := types.BankKeeper(app.BankKeeper)
	ps := paramtypes.Subspace{}
	result := keeper.NewKeeper(cdc, key, bankkeeper, accountKeeper, ps)

	res := result.Codec()
	t.Log(res)
}

func TestKeeper_getAccountKeeper(t *testing.T) {
	app, ctx := test.CreateTestApp(false)
	t.Log(ctx)
	key := sdk.NewKVStoreKey("rowan")
	cdc := codec.BinaryCodec(app.AppCodec())
	accountKeeper := types.AccountKeeper(app.AccountKeeper)
	bankkeeper := types.BankKeeper(app.BankKeeper)
	ps := paramtypes.Subspace{}
	result := keeper.NewKeeper(cdc, key, bankkeeper, accountKeeper, ps)

	res := result.GetAccountKeeper()
	t.Log(res)

}

func TestKeeper_hasCoins(t *testing.T) {
	app, ctx := test.CreateTestApp(false)

	key := sdk.NewKVStoreKey("rowan")
	cdc := codec.BinaryCodec(app.AppCodec())
	accountKeeper := types.AccountKeeper(app.AccountKeeper)
	bankkeeper := types.BankKeeper(app.BankKeeper)
	ps := paramtypes.Subspace{}
	result := keeper.NewKeeper(cdc, key, bankkeeper, accountKeeper, ps)
	user := sdk.AccAddress("xyz")
	coins := sdk.Coins{}
	res := result.HasCoins(ctx, user, coins)
	t.Log(res)
}
