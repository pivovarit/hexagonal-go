package controller

import "github.com/pivovarit/hexagonal-go/app/account"

type AccountController struct {
	accounts *account.Service
}

func NewAccountController(accounts *account.Service) *AccountController {
	return &AccountController{accounts: accounts}
}
