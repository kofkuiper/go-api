package handlers

import (
	"net/http"

	"github.com/kofkuiper/echo-api/services"
	"github.com/labstack/echo/v4"
)

type plutoHandler struct {
	plutoSrv services.PlutoService
}

func NewPlutoHandler(plutoSrv services.PlutoService) plutoHandler {
	return plutoHandler{plutoSrv}
}

func (p plutoHandler) Info(c echo.Context) error {
	info, err := p.plutoSrv.ChainInfo()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{"info": info})
}
