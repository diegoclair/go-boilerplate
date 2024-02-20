package accountroute

import (
	"sync"

	"github.com/diegoclair/go_boilerplate/domain/contract"
	"github.com/diegoclair/go_boilerplate/domain/entity"
	"github.com/diegoclair/go_boilerplate/transport/rest/routeutils"
	"github.com/diegoclair/go_boilerplate/transport/rest/viewmodel"
	"github.com/diegoclair/go_utils-lib/v2/validator"

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

type Controller struct {
	accountService contract.AccountService
	utils          routeutils.Utils
	validator      validator.Validator
}

func NewController(accountService contract.AccountService, utils routeutils.Utils, validator validator.Validator) *Controller {
	once.Do(func() {
		instance = &Controller{
			accountService: accountService,
			utils:          utils,
			validator:      validator,
		}
	})

	return instance
}

func (s *Controller) handleAddAccount(c echo.Context) error {
	ctx := s.utils.Req().GetContext(c)

	input := viewmodel.AddAccount{}
	err := c.Bind(&input)
	if err != nil {
		return s.utils.Resp().ResponseBadRequestError(c, err)
	}

	err = input.Validate(s.validator)
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	account := entity.Account{
		Name:     input.Name,
		CPF:      input.CPF,
		Password: input.Password,
	}

	err = s.accountService.CreateAccount(ctx, account)
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	return s.utils.Resp().ResponseCreated(c)
}

func (s *Controller) handleAddBalance(c echo.Context) error {
	ctx := s.utils.Req().GetContext(c)

	input := viewmodel.AddBalance{}
	err := c.Bind(&input)
	if err != nil {
		return s.utils.Resp().ResponseBadRequestError(c, err)
	}

	err = input.Validate(s.validator)
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	accountUUID, err := s.utils.Req().GetAndValidateParam(c, "account_uuid", "Invalid account_uuid")
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	err = s.accountService.AddBalance(ctx, accountUUID, input.Amount)
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	return s.utils.Resp().ResponseCreated(c)
}

func (s *Controller) handleGetAccounts(c echo.Context) error {
	ctx := s.utils.Req().GetContext(c)

	take, skip := s.utils.Req().GetPagingParams(c, "page", "quantity")

	accounts, totalRecords, err := s.accountService.GetAccounts(ctx, take, skip)
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	response := []viewmodel.Account{}
	for _, account := range accounts {
		item := viewmodel.Account{}
		item.FillFromEntity(account)
		response = append(response, item)
	}

	responsePaginated := routeutils.BuildPaginatedResult(response, skip, take, totalRecords)

	return s.utils.Resp().ResponseAPIOk(c, responsePaginated)
}

func (s *Controller) handleGetAccountByID(c echo.Context) error {
	ctx := s.utils.Req().GetContext(c)

	accountUUID, err := s.utils.Req().GetAndValidateParam(c, "account_uuid", "Invalid account_uuid")
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	account, err := s.accountService.GetAccountByUUID(ctx, accountUUID)
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	response := viewmodel.Account{}
	response.FillFromEntity(account)

	return s.utils.Resp().ResponseAPIOk(c, response)
}

func (s *Controller) handleGetAccountBalanceByID(c echo.Context) error {
	ctx := s.utils.Req().GetContext(c)

	accountUUID, err := s.utils.Req().GetAndValidateParam(c, "account_id", "Invalid account_id")
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	account, err := s.accountService.GetAccountByUUID(ctx, accountUUID)
	if err != nil {
		return s.utils.Resp().HandleAPIError(c, err)
	}

	response := viewmodel.Account{
		Balance: account.Balance,
	}

	return s.utils.Resp().ResponseAPIOk(c, response)
}