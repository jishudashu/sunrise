// Code generated by MockGen. DO NOT EDIT.
// Source: x/swap/types/expected_keepers.go
//
// Generated by this command:
//
//	mockgen -source=x/swap/types/expected_keepers.go -destination=x/swap/testutil/expected_keepers_mocks.go -package=testutil
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/ibc-go/v10/modules/apps/transfer/types"
	types1 "github.com/sunriselayer/sunrise/x/liquiditypool/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
	isgomock struct{}
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
func (m *MockAccountKeeper) GetAccount(arg0 context.Context, arg1 types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), arg0, arg1)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
	isgomock struct{}
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

// IsSendEnabledCoins mocks base method.
func (m *MockBankKeeper) IsSendEnabledCoins(ctx context.Context, coins ...types.Coin) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range coins {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsSendEnabledCoins", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsSendEnabledCoins indicates an expected call of IsSendEnabledCoins.
func (mr *MockBankKeeperMockRecorder) IsSendEnabledCoins(ctx any, coins ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, coins...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSendEnabledCoins", reflect.TypeOf((*MockBankKeeper)(nil).IsSendEnabledCoins), varargs...)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx context.Context, fromAddr, toAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, fromAddr, toAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, fromAddr, toAddr, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(arg0 context.Context, arg1 types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", arg0, arg1)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), arg0, arg1)
}

// MockTransferKeeper is a mock of TransferKeeper interface.
type MockTransferKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockTransferKeeperMockRecorder
	isgomock struct{}
}

// MockTransferKeeperMockRecorder is the mock recorder for MockTransferKeeper.
type MockTransferKeeperMockRecorder struct {
	mock *MockTransferKeeper
}

// NewMockTransferKeeper creates a new mock instance.
func NewMockTransferKeeper(ctrl *gomock.Controller) *MockTransferKeeper {
	mock := &MockTransferKeeper{ctrl: ctrl}
	mock.recorder = &MockTransferKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferKeeper) EXPECT() *MockTransferKeeperMockRecorder {
	return m.recorder
}

// Transfer mocks base method.
func (m *MockTransferKeeper) Transfer(ctx context.Context, msg *types0.MsgTransfer) (*types0.MsgTransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, msg)
	ret0, _ := ret[0].(*types0.MsgTransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockTransferKeeperMockRecorder) Transfer(ctx, msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockTransferKeeper)(nil).Transfer), ctx, msg)
}

// MockLiquidityPoolKeeper is a mock of LiquidityPoolKeeper interface.
type MockLiquidityPoolKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockLiquidityPoolKeeperMockRecorder
	isgomock struct{}
}

// MockLiquidityPoolKeeperMockRecorder is the mock recorder for MockLiquidityPoolKeeper.
type MockLiquidityPoolKeeperMockRecorder struct {
	mock *MockLiquidityPoolKeeper
}

// NewMockLiquidityPoolKeeper creates a new mock instance.
func NewMockLiquidityPoolKeeper(ctrl *gomock.Controller) *MockLiquidityPoolKeeper {
	mock := &MockLiquidityPoolKeeper{ctrl: ctrl}
	mock.recorder = &MockLiquidityPoolKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiquidityPoolKeeper) EXPECT() *MockLiquidityPoolKeeperMockRecorder {
	return m.recorder
}

// CalculateResultExactAmountIn mocks base method.
func (m *MockLiquidityPoolKeeper) CalculateResultExactAmountIn(ctx types.Context, pool types1.Pool, tokenIn types.Coin, denomOut string, feeEnabled bool) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateResultExactAmountIn", ctx, pool, tokenIn, denomOut, feeEnabled)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateResultExactAmountIn indicates an expected call of CalculateResultExactAmountIn.
func (mr *MockLiquidityPoolKeeperMockRecorder) CalculateResultExactAmountIn(ctx, pool, tokenIn, denomOut, feeEnabled any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateResultExactAmountIn", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).CalculateResultExactAmountIn), ctx, pool, tokenIn, denomOut, feeEnabled)
}

// CalculateResultExactAmountOut mocks base method.
func (m *MockLiquidityPoolKeeper) CalculateResultExactAmountOut(ctx types.Context, pool types1.Pool, tokenOut types.Coin, denomIn string, feeEnabled bool) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateResultExactAmountOut", ctx, pool, tokenOut, denomIn, feeEnabled)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateResultExactAmountOut indicates an expected call of CalculateResultExactAmountOut.
func (mr *MockLiquidityPoolKeeperMockRecorder) CalculateResultExactAmountOut(ctx, pool, tokenOut, denomIn, feeEnabled any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateResultExactAmountOut", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).CalculateResultExactAmountOut), ctx, pool, tokenOut, denomIn, feeEnabled)
}

// GetPool mocks base method.
func (m *MockLiquidityPoolKeeper) GetPool(ctx context.Context, id uint64) (types1.Pool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPool", ctx, id)
	ret0, _ := ret[0].(types1.Pool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPool indicates an expected call of GetPool.
func (mr *MockLiquidityPoolKeeperMockRecorder) GetPool(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPool", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).GetPool), ctx, id)
}

// SwapExactAmountIn mocks base method.
func (m *MockLiquidityPoolKeeper) SwapExactAmountIn(ctx types.Context, sender types.AccAddress, pool types1.Pool, tokenIn types.Coin, denomOut string, feeEnabled bool) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwapExactAmountIn", ctx, sender, pool, tokenIn, denomOut, feeEnabled)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SwapExactAmountIn indicates an expected call of SwapExactAmountIn.
func (mr *MockLiquidityPoolKeeperMockRecorder) SwapExactAmountIn(ctx, sender, pool, tokenIn, denomOut, feeEnabled any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwapExactAmountIn", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).SwapExactAmountIn), ctx, sender, pool, tokenIn, denomOut, feeEnabled)
}

// SwapExactAmountOut mocks base method.
func (m *MockLiquidityPoolKeeper) SwapExactAmountOut(ctx types.Context, sender types.AccAddress, pool types1.Pool, tokenOut types.Coin, denomIn string, feeEnabled bool) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwapExactAmountOut", ctx, sender, pool, tokenOut, denomIn, feeEnabled)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SwapExactAmountOut indicates an expected call of SwapExactAmountOut.
func (mr *MockLiquidityPoolKeeperMockRecorder) SwapExactAmountOut(ctx, sender, pool, tokenOut, denomIn, feeEnabled any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwapExactAmountOut", reflect.TypeOf((*MockLiquidityPoolKeeper)(nil).SwapExactAmountOut), ctx, sender, pool, tokenOut, denomIn, feeEnabled)
}
