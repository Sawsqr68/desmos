package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		params types.Params
		check  func(ctx sdk.Context)
	}{
		{
			name:   "default params are saved correctly",
			params: types.DefaultParams(),
			check: func(ctx sdk.Context) {
				stored := suite.k.GetParams(ctx)
				suite.Require().Equal(stored, types.DefaultParams())
			},
		},
		{
			name: "params are overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			params: types.NewParams(sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(100)))),
			check: func(ctx sdk.Context) {
				stored := suite.k.GetParams(ctx)
				suite.Require().Equal(types.NewParams(sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(100)))), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.SetParams(ctx, tc.params)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expParams types.Params
	}{
		{
			name: "params are returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(100)))))
			},
			expParams: types.NewParams(sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(100)))),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			params := suite.k.GetParams(ctx)
			suite.Require().Equal(tc.expParams, params)
		})
	}
}
