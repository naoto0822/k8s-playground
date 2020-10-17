package pkg

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Ruok(c echo.Context) error {
	msg := "imok"
	return c.String(http.StatusOK, msg)
}

func (h Handler) GetSobaList(c echo.Context) error {
	return nil
}
