package keeper_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/clp/keeper"
	"github.com/Sifchain/sifnode/x/clp/test"
	"github.com/tendermint/tendermint/abci/types"
)

func Test_BeginBlock(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	clpKeeper.BeginBlock(ctx)

}

func Test_EndBlock(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	res := keeper.EndBlock(ctx, types.RequestEndBlock{}, clpKeeper)
	t.Log(res)
}

func Test_GetCurrentRewardPeriod(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	params := clpKeeper.GetParams(ctx)
	res := clpKeeper.GetCurrentRewardPeriod(ctx, params)
	t.Log(res)

}

func Test_PruneRewardPeriods(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	params := clpKeeper.GetParams(ctx)

	clpKeeper.PruneRewardPeriods(ctx, params)

}
