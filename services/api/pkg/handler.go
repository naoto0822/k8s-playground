package pkg

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	processorService *ProcessorService
}

func NewHandler(service *ProcessorService) *Handler {
	return &Handler{
		processorService: service,
	}
}

func (h *Handler) Ruok(c echo.Context) error {
	msg := "imok"
	return c.String(http.StatusOK, msg)
}

func (h *Handler) GetSobaList(c echo.Context) error {
	ctx := context.Background()
	sobas, err := h.processorService.GetSobaList(ctx)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, sobas)
}
