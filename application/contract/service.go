package contract

import (
	"context"

	"github.com/diegoclair/go_boilerplate/application/dto"
	"github.com/diegoclair/go_boilerplate/domain/entity"
)

type AccountService interface {
	CreateAccount(ctx context.Context, account entity.Account) (err error)
	AddBalance(ctx context.Context, accountUUID string, amount float64) (err error)
	GetAccounts(ctx context.Context, take, skip int64) (accounts []entity.Account, totalRecords int64, err error)
	GetAccountByUUID(ctx context.Context, accountUUID string) (account entity.Account, err error)
}

type AuthService interface {
	Login(ctx context.Context, cpf, secret string) (account entity.Account, err error)
	CreateSession(ctx context.Context, session dto.Session) (err error)
	GetSessionByUUID(ctx context.Context, sessionUUID string) (session dto.Session, err error)
}

type TransferService interface {
	CreateTransfer(ctx context.Context, transfer entity.Transfer) (err error)
	GetTransfers(ctx context.Context, take, skip int64) (transfers []entity.Transfer, totalRecords int64, err error)
}
