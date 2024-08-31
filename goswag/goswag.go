package main

import (
	_ "github.com/diegoclair/go_boilerplate/transport/rest/viewmodel"
)

//	@Summary		Logout
//	@Description	Logout the user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user-token	header	string	true	"User access token"
//	@Success		200
//	@Router			/auth/logout [post]
func handleLogout() {} //nolint:unused 

//	@Summary		Add a new transfer
//	@Description	Add a new transfer
//	@Tags			transfers
//	@Accept			json
//	@Produce		json
//	@Param			request		body	viewmodel.TransferReq	true	"Request"
//	@Param			user-token	header	string					true	"User access token"
//	@Success		201
//	@Router			/transfers [post]
func handleAddTransfer() {} //nolint:unused 

//	@Summary		Get all transfers
//	@Description	Get all transfers with paginated response
//	@Tags			transfers
//	@Produce		json
//	@Param			user-token	header		string	true	"User access token"
//	@Success		200			{object}	viewmodel.PaginatedResponse[[]viewmodel.TransferResp]
//	@Router			/transfers [get]
func handleGetTransfers() {} //nolint:unused 

//	@Summary		Add a new account
//	@Description	Add a new account
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			request	body	viewmodel.AddAccount	true	"Request"
//	@Success		201
//	@Router			/accounts [post]
func handleAddAccount() {} //nolint:unused 

//	@Summary		Add balance to an account
//	@Description	Add balance to an account by account_uuid
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			request			body	viewmodel.AddBalance	true	"Request"
//	@Param			account_uuid	path	string					true	"account uuid"
//	@Success		201
//	@Router			/accounts/:account_uuid/balance [post]
func handleAddBalance() {} //nolint:unused 

//	@Summary		Get all accounts
//	@Description	Get all accounts with paginated response
//	@Tags			accounts
//	@Produce		json
//	@Param			page		query		string	false	"number of page you want"
//	@Param			quantity	query		string	false	"quantity of items per page"
//	@Success		200			{object}	viewmodel.PaginatedResponse[[]viewmodel.AccountResponse]
//	@Router			/accounts [get]
func handleGetAccounts() {} //nolint:unused 

//	@Summary		Get account by ID
//	@Description	Get account by it UUID value
//	@Tags			accounts
//	@Produce		json
//	@Param			account_uuid	path		string	true	"account uuid"
//	@Success		200				{object}	viewmodel.AccountResponse
//	@Router			/accounts/:account_uuid/ [get]
func handleGetAccountByID() {} //nolint:unused 

//	@Summary		Login
//	@Description	Login
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		viewmodel.Login	true	"Request"
//	@Success		200		{object}	viewmodel.LoginResponse
//	@Router			/auth/login [post]
func handleLogin() {} //nolint:unused 

//	@Summary		Refresh Token
//	@Description	Generate a new token using the refresh token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		viewmodel.RefreshTokenRequest	true	"Request"
//	@Success		200		{object}	viewmodel.RefreshTokenResponse
//	@Router			/auth/refresh-token [post]
func handleRefreshToken() {} //nolint:unused 

//	@Summary		Ping the server
//	@Description	Ping the server to check if it is alive
//	@Tags			ping
//	@Produce		json
//	@Success		200	{object}	pingroute.pingResponse
//	@Router			/ping/ [get]
func handlePing() {} //nolint:unused 

