package service

import (
	"testing"

	"github.com/diegoclair/go_boilerplate/domain/contract"
	"github.com/diegoclair/go_boilerplate/infra/logger"
	"github.com/diegoclair/go_boilerplate/mocks"
	"github.com/diegoclair/go_boilerplate/util/config"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type repoMock struct {
	mockAuthRepo     *mocks.MockAuthRepo
	mockAccountRepo  *mocks.MockAccountRepo
	mockCacheManager *mocks.MockCacheManager
}

func newServiceTestMock(t *testing.T) (repoMocks repoMock, svc *Service, ctrl *gomock.Controller) {

	cfg, err := config.GetConfigEnvironment("../../" + config.ConfigDefaultFilepath)
	require.NoError(t, err)

	ctrl = gomock.NewController(t)
	log := logger.New(*cfg)

	repoMocks = repoMock{
		mockAccountRepo:  mocks.NewMockAccountRepo(ctrl),
		mockCacheManager: mocks.NewMockCacheManager(ctrl),
		mockAuthRepo:     mocks.NewMockAuthRepo(ctrl),
	}

	dataManagerMock := newDataMock(ctrl, repoMocks)

	svc = New(dataManagerMock, cfg, repoMocks.mockCacheManager, log)

	return
}

type dataMock struct {
	ctrl  *gomock.Controller
	mocks repoMock
}

func newDataMock(ctrl *gomock.Controller, mocks repoMock) contract.DataManager {
	return &dataMock{
		ctrl:  ctrl,
		mocks: mocks,
	}
}

func (d *dataMock) Begin() (contract.Transaction, error) {
	return mocks.NewMockTransaction(d.ctrl), nil
}

func (d *dataMock) Account() contract.AccountRepo {
	return d.mocks.mockAccountRepo
}

func (d *dataMock) Auth() contract.AuthRepo {
	return d.mocks.mockAuthRepo
}
