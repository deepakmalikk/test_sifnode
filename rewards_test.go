package keeper_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/clp/keeper"
	"github.com/Sifchain/sifnode/x/clp/test"
	"github.com/Sifchain/sifnode/x/clp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func Test_BeginBlock(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	clpKeeper.BeginBlock(ctx)

}

func Test_EndBlock(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	res := keeper.EndBlock(ctx, abci.RequestEndBlock{}, clpKeeper)
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

// func Test_DistributeDepthRewards(t *testing.T) {
// 	ctx, app := test.CreateTestAppClp(false)
// 	clpKeeper := app.ClpKeeper
// 	peroids := types.RewardPeriod{}
// 	peroid := &peroids
// 	pool := []*types.Pool{}
// 	res := clpKeeper.DistributeDepthRewards(ctx, peroid, pool)
// 	t.Log(res)
// }

func Test_GetRewardExecution(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	res := clpKeeper.GetRewardExecution(ctx)
	t.Log(res)

}

func Test_SetRewardExecution(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	reward := types.RewardExecution{}
	clpKeeper.SetRewardExecution(ctx, reward)

}

func Test_UseUnlockedLiquidity(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	lp := test.GenerateRandomLP(1)[0]
	clpKeeper.SetLiquidityProvider(ctx, &lp)
	units := sdk.NewUint(1000)
	res := clpKeeper.UseUnlockedLiquidity(ctx, lp, units)
	t.Log(res)
	lp1 := test.GenerateRandomLP(1)[0]
	clpKeeper.SetLiquidityProvider(ctx, &lp)
	units1 := sdk.NewUint(0)
	res1 := clpKeeper.UseUnlockedLiquidity(ctx, lp1, units1)
	t.Log(res1)

}

func Test_PruneUnlockRecords(t *testing.T) {
	ctx, app := test.CreateTestAppClp(false)
	clpKeeper := app.ClpKeeper
	lp := test.GenerateRandomLP(1)[0]
	clpKeeper.SetLiquidityProvider(ctx, &lp)
	lockPeriod := sdk.NewUint(1000)
	cancelPeriod := sdk.NewUint(10000)
	clpKeeper.PruneUnlockRecords(ctx, lp, lockPeriod.Uint64(), cancelPeriod.Uint64())

}
