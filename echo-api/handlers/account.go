package handlers

import (
	"net/http"
	"strings"

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
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, account)
}

func (a accountHandler) Login(c echo.Context) error {
	body := new(services.LoginRequest)
	if err := c.Bind(body); err != nil {
		return err
	}

	errors := services.Validate(body)
	if errors != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"errors": errors})
	}

	token, err := a.accountSrv.Login(*body)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (a accountHandler) Validate(c echo.Context) error {
	authorization := c.Request().Header["Authorization"]
	if len(authorization) == 0 {
		return c.String(http.StatusUnauthorized, "unauthorized")

	}
	bearerToken := authorization[0]
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	claims, err := a.accountSrv.Validate(token)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"claims": claims})
}
