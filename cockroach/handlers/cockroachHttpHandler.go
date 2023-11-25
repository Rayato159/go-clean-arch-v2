package handlers

import (
	"net/http"

	"github.com/Rayato159/go-clean-arch-v2/cockroach/models"
	"github.com/Rayato159/go-clean-arch-v2/cockroach/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type cockroachHttpHandler struct {
	cockroachUsecase usecases.CockroachUsecase
}

func NewCockroachHttpHandler(cockroachUsecase usecases.CockroachUsecase) CockroachHandler {
	return &cockroachHttpHandler{
		cockroachUsecase: cockroachUsecase,
	}
}

func (h *cockroachHttpHandler) DetectCockroach(c echo.Context) error {
	reqBody := new(models.AddCockroachData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}

	if err := h.cockroachUsecase.CockroachDataProcessing(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing data failed")
	}

	return response(c, http.StatusOK, "Success 🪳🪳🪳")
}
