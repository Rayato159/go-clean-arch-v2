package handlers

import (
	"github.com/Rayato159/go-clean-arch-v2/cockroach/usecases"

	"github.com/labstack/echo/v4"
)

type cockroachHttpHandler struct {
	cockroachUsecase usecases.CockroachUsecase
}

func (h *cockroachHttpHandler) DetectCockroach(c echo.Context) error {
	return nil
}
