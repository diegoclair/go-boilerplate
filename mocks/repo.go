// Code generated by MockGen. DO NOT EDIT.
// Source: application/contract/repo.go
//
// Generated by this command:
//
//	mockgen -package mocks -source=application/contract/repo.go -destination=mocks/repo.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	contract "github.com/diegoclair/go_boilerplate/domain/contract"
	dto "github.com/diegoclair/go_boilerplate/application/dto"
	entity "github.com/diegoclair/go_boilerplate/domain/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockDataManager is a mock of DataManager interface.
type MockDataManager struct {
	ctrl     *gomock.Controller
	recorder *MockDataManagerMockRecorder
}

// MockDataManagerMockRecorder is the mock recorder for MockDataManager.
type MockDataManagerMockRecorder struct {
	mock *MockDataManager
}

// NewMockDataManager creates a new mock instance.
func NewMockDataManager(ctrl *gomock.Controller) *MockDataManager {
	mock := &MockDataManager{ctrl: ctrl}
	mock.recorder = &MockDataManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataManager) EXPECT() *MockDataManagerMockRecorder {
	return m.recorder
}

// Account mocks base method.
func (m *MockDataManager) Account() contract.AccountRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Account")
	ret0, _ := ret[0].(contract.AccountRepo)
	return ret0
}

// Account indicates an expected call of Account.
func (mr *MockDataManagerMockRecorder) Account() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Account", reflect.TypeOf((*MockDataManager)(nil).Account))
}

// Auth mocks base method.
func (m *MockDataManager) Auth() contract.AuthRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth")
	ret0, _ := ret[0].(contract.AuthRepo)
	return ret0
}

// Auth indicates an expected call of Auth.
func (mr *MockDataManagerMockRecorder) Auth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockDataManager)(nil).Auth))
}

// WithTransaction mocks base method.
func (m *MockDataManager) WithTransaction(ctx context.Context, fn func(contract.DataManager) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTransaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTransaction indicates an expected call of WithTransaction.
func (mr *MockDataManagerMockRecorder) WithTransaction(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransaction", reflect.TypeOf((*MockDataManager)(nil).WithTransaction), ctx, fn)
}

// MockAuthRepo is a mock of AuthRepo interface.
type MockAuthRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepoMockRecorder
}

// MockAuthRepoMockRecorder is the mock recorder for MockAuthRepo.
type MockAuthRepoMockRecorder struct {
	mock *MockAuthRepo
}

// NewMockAuthRepo creates a new mock instance.
func NewMockAuthRepo(ctrl *gomock.Controller) *MockAuthRepo {
	mock := &MockAuthRepo{ctrl: ctrl}
	mock.recorder = &MockAuthRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepo) EXPECT() *MockAuthRepoMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockAuthRepo) CreateSession(ctx context.Context, session dto.Session) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, session)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockAuthRepoMockRecorder) CreateSession(ctx, session any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockAuthRepo)(nil).CreateSession), ctx, session)
}

// GetSessionByUUID mocks base method.
func (m *MockAuthRepo) GetSessionByUUID(ctx context.Context, sessionUUID string) (dto.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByUUID", ctx, sessionUUID)
	ret0, _ := ret[0].(dto.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByUUID indicates an expected call of GetSessionByUUID.
func (mr *MockAuthRepoMockRecorder) GetSessionByUUID(ctx, sessionUUID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByUUID", reflect.TypeOf((*MockAuthRepo)(nil).GetSessionByUUID), ctx, sessionUUID)
}

// MockAccountRepo is a mock of AccountRepo interface.
type MockAccountRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepoMockRecorder
}

// MockAccountRepoMockRecorder is the mock recorder for MockAccountRepo.
type MockAccountRepoMockRecorder struct {
	mock *MockAccountRepo
}

// NewMockAccountRepo creates a new mock instance.
func NewMockAccountRepo(ctrl *gomock.Controller) *MockAccountRepo {
	mock := &MockAccountRepo{ctrl: ctrl}
	mock.recorder = &MockAccountRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepo) EXPECT() *MockAccountRepoMockRecorder {
	return m.recorder
}

// AddTransfer mocks base method.
func (m *MockAccountRepo) AddTransfer(ctx context.Context, transferUUID string, accountOriginID, accountDestinationID int64, amount float64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransfer", ctx, transferUUID, accountOriginID, accountDestinationID, amount)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTransfer indicates an expected call of AddTransfer.
func (mr *MockAccountRepoMockRecorder) AddTransfer(ctx, transferUUID, accountOriginID, accountDestinationID, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransfer", reflect.TypeOf((*MockAccountRepo)(nil).AddTransfer), ctx, transferUUID, accountOriginID, accountDestinationID, amount)
}

// CreateAccount mocks base method.
func (m *MockAccountRepo) CreateAccount(ctx context.Context, account entity.Account) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, account)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountRepoMockRecorder) CreateAccount(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountRepo)(nil).CreateAccount), ctx, account)
}

// GetAccountByDocument mocks base method.
func (m *MockAccountRepo) GetAccountByDocument(ctx context.Context, encryptedCPF string) (entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByDocument", ctx, encryptedCPF)
	ret0, _ := ret[0].(entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByDocument indicates an expected call of GetAccountByDocument.
func (mr *MockAccountRepoMockRecorder) GetAccountByDocument(ctx, encryptedCPF any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByDocument", reflect.TypeOf((*MockAccountRepo)(nil).GetAccountByDocument), ctx, encryptedCPF)
}

// GetAccountByUUID mocks base method.
func (m *MockAccountRepo) GetAccountByUUID(ctx context.Context, accountUUID string) (entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByUUID", ctx, accountUUID)
	ret0, _ := ret[0].(entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByUUID indicates an expected call of GetAccountByUUID.
func (mr *MockAccountRepoMockRecorder) GetAccountByUUID(ctx, accountUUID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByUUID", reflect.TypeOf((*MockAccountRepo)(nil).GetAccountByUUID), ctx, accountUUID)
}

// GetAccounts mocks base method.
func (m *MockAccountRepo) GetAccounts(ctx context.Context, take, skip int64) ([]entity.Account, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", ctx, take, skip)
	ret0, _ := ret[0].([]entity.Account)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockAccountRepoMockRecorder) GetAccounts(ctx, take, skip any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockAccountRepo)(nil).GetAccounts), ctx, take, skip)
}

// GetTransfersByAccountID mocks base method.
func (m *MockAccountRepo) GetTransfersByAccountID(ctx context.Context, accountID, take, skip int64, origin bool) ([]entity.Transfer, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfersByAccountID", ctx, accountID, take, skip, origin)
	ret0, _ := ret[0].([]entity.Transfer)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransfersByAccountID indicates an expected call of GetTransfersByAccountID.
func (mr *MockAccountRepoMockRecorder) GetTransfersByAccountID(ctx, accountID, take, skip, origin any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfersByAccountID", reflect.TypeOf((*MockAccountRepo)(nil).GetTransfersByAccountID), ctx, accountID, take, skip, origin)
}

// UpdateAccountBalance mocks base method.
func (m *MockAccountRepo) UpdateAccountBalance(ctx context.Context, accountID int64, balance float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountBalance", ctx, accountID, balance)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccountBalance indicates an expected call of UpdateAccountBalance.
func (mr *MockAccountRepoMockRecorder) UpdateAccountBalance(ctx, accountID, balance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountBalance", reflect.TypeOf((*MockAccountRepo)(nil).UpdateAccountBalance), ctx, accountID, balance)
}
