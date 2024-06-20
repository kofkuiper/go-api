package handlers

import (
	"net/http"

	"github.com/kofkuiper/echo-api/services"
	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	accountSrv services.AccountService
}

func NewAccountHandler(accountSrv services.AccountService) accountHandler {
	return accountHandler{accountSrv: accountSrv}
}

func (a accountHandler) SignUp(c echo.Context) error {
	body := new(services.SignUpRequest)
	if err := c.Bind(body); err != nil {
		return err
	}

	errors := services.Validate(body)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"errors": errors})
	}

	account, err := a.accountSrv.SignUp(*body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "unexpected error")
	}

	return c.JSON(http.StatusCreated, account)
}
