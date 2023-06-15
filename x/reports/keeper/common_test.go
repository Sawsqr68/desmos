package keeper_test

import (
	"testing"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	profileskeeper "github.com/desmos-labs/desmos/v5/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v5/x/profiles/types"

	postskeeper "github.com/desmos-labs/desmos/v5/x/posts/keeper"
	poststypes "github.com/desmos-labs/desmos/v5/x/posts/types"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v5/app"
	relationshipskeeper "github.com/desmos-labs/desmos/v5/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v5/x/relationships/types"
	"github.com/desmos-labs/desmos/v5/x/reports/keeper"
	"github.com/desmos-labs/desmos/v5/x/reports/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v5/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context

	storeKey storetypes.StoreKey
	k        keeper.Keeper

	ak profileskeeper.Keeper
	sk subspaceskeeper.Keeper
	rk relationshipskeeper.Keeper
	pk postskeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(
		paramstypes.StoreKey, authtypes.StoreKey,
		profilestypes.StoreKey, relationshipstypes.StoreKey,
		subspacestypes.StoreKey, poststypes.StoreKey,
		types.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}
	for _, tKey := range tKeys {
		ms.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Define keeper
	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey], nil, nil, authtypes.NewModuleAddress("gov").String())
	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, keys[relationshipstypes.StoreKey], suite.sk)
	authKeeper := authkeeper.NewAccountKeeper(suite.cdc, keys[authtypes.StoreKey], authtypes.ProtoBaseAccount, app.GetMaccPerms(), "cosmos", authtypes.NewModuleAddress("gov").String())
	suite.ak = profileskeeper.NewKeeper(suite.cdc, suite.legacyAminoCdc, keys[profilestypes.StoreKey], authKeeper, suite.rk, nil, nil, nil, authtypes.NewModuleAddress("gov").String())
	suite.pk = postskeeper.NewKeeper(suite.cdc, keys[poststypes.StoreKey], suite.ak, suite.sk, suite.rk, authtypes.NewModuleAddress("gov").String())
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		suite.ak,
		suite.sk,
		suite.rk,
		suite.pk,
		authtypes.NewModuleAddress("gov").String(),
	)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
