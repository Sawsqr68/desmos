// Code generated by MockGen. DO NOT EDIT.
// Source: ./x/tokenfactory/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/bank/types"
	gomock "github.com/golang/mock/gomock"

	types2 "github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx types.Context, moduleName string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, moduleName, amt)
}

// GetDenomMetaData mocks base method.
func (m *MockBankKeeper) GetDenomMetaData(ctx types.Context, denom string) (types1.Metadata, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDenomMetaData", ctx, denom)
	ret0, _ := ret[0].(types1.Metadata)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDenomMetaData indicates an expected call of GetDenomMetaData.
func (mr *MockBankKeeperMockRecorder) GetDenomMetaData(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).GetDenomMetaData), ctx, denom)
}

// HasBalance mocks base method.
func (m *MockBankKeeper) HasBalance(ctx types.Context, addr types.AccAddress, amt types.Coin) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBalance", ctx, addr, amt)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBalance indicates an expected call of HasBalance.
func (mr *MockBankKeeperMockRecorder) HasBalance(ctx, addr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBalance", reflect.TypeOf((*MockBankKeeper)(nil).HasBalance), ctx, addr, amt)
}

// HasSupply mocks base method.
func (m *MockBankKeeper) HasSupply(ctx types.Context, denom string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSupply", ctx, denom)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSupply indicates an expected call of HasSupply.
func (mr *MockBankKeeperMockRecorder) HasSupply(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSupply", reflect.TypeOf((*MockBankKeeper)(nil).HasSupply), ctx, denom)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx types.Context, moduleName string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx types.Context, fromAddr, toAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, fromAddr, toAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, fromAddr, toAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SetDenomMetaData mocks base method.
func (m *MockBankKeeper) SetDenomMetaData(ctx types.Context, denomMetaData types1.Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDenomMetaData", ctx, denomMetaData)
}

// SetDenomMetaData indicates an expected call of SetDenomMetaData.
func (mr *MockBankKeeperMockRecorder) SetDenomMetaData(ctx, denomMetaData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDenomMetaData", reflect.TypeOf((*MockBankKeeper)(nil).SetDenomMetaData), ctx, denomMetaData)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(arg0 types.Context, arg1 types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), arg0, arg1)
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx types.Context, moduleName string) types0.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, moduleName)
	ret0, _ := ret[0].(types0.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, moduleName)
}

// MockSubspacesKeeper is a mock of SubspacesKeeper interface.
type MockSubspacesKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSubspacesKeeperMockRecorder
}

// MockSubspacesKeeperMockRecorder is the mock recorder for MockSubspacesKeeper.
type MockSubspacesKeeperMockRecorder struct {
	mock *MockSubspacesKeeper
}

// NewMockSubspacesKeeper creates a new mock instance.
func NewMockSubspacesKeeper(ctrl *gomock.Controller) *MockSubspacesKeeper {
	mock := &MockSubspacesKeeper{ctrl: ctrl}
	mock.recorder = &MockSubspacesKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubspacesKeeper) EXPECT() *MockSubspacesKeeperMockRecorder {
	return m.recorder
}

// GetAllSubspaces mocks base method.
func (m *MockSubspacesKeeper) GetAllSubspaces(ctx types.Context) []types2.Subspace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubspaces", ctx)
	ret0, _ := ret[0].([]types2.Subspace)
	return ret0
}

// GetAllSubspaces indicates an expected call of GetAllSubspaces.
func (mr *MockSubspacesKeeperMockRecorder) GetAllSubspaces(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubspaces", reflect.TypeOf((*MockSubspacesKeeper)(nil).GetAllSubspaces), ctx)
}

// GetSubspace mocks base method.
func (m *MockSubspacesKeeper) GetSubspace(ctx types.Context, subspaceID uint64) (types2.Subspace, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubspace", ctx, subspaceID)
	ret0, _ := ret[0].(types2.Subspace)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetSubspace indicates an expected call of GetSubspace.
func (mr *MockSubspacesKeeperMockRecorder) GetSubspace(ctx, subspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubspace", reflect.TypeOf((*MockSubspacesKeeper)(nil).GetSubspace), ctx, subspaceID)
}

// GetUsersWithRootPermissions mocks base method.
func (m *MockSubspacesKeeper) GetUsersWithRootPermissions(ctx types.Context, subspaceID uint64, permission types2.Permissions) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersWithRootPermissions", ctx, subspaceID, permission)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetUsersWithRootPermissions indicates an expected call of GetUsersWithRootPermissions.
func (mr *MockSubspacesKeeperMockRecorder) GetUsersWithRootPermissions(ctx, subspaceID, permission interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersWithRootPermissions", reflect.TypeOf((*MockSubspacesKeeper)(nil).GetUsersWithRootPermissions), ctx, subspaceID, permission)
}

// HasPermission mocks base method.
func (m *MockSubspacesKeeper) HasPermission(ctx types.Context, subspaceID uint64, sectionID uint32, user string, permission types2.Permission) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermission", ctx, subspaceID, sectionID, user, permission)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPermission indicates an expected call of HasPermission.
func (mr *MockSubspacesKeeperMockRecorder) HasPermission(ctx, subspaceID, sectionID, user, permission interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermission", reflect.TypeOf((*MockSubspacesKeeper)(nil).HasPermission), ctx, subspaceID, sectionID, user, permission)
}
