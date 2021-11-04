package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mdmims/go-echo-crud/repository"
	"net/http"
)

// ping godoc
// @Summary server health check
// @Accept json
// @Produce json
// @Success 200 {object} repository.healthResponse
// @Failure 400 {object} repository.Error
// @Failure 404 {object} repository.Error
// @Failure 500 {object} repository.Error
// @Router /health/ping [get]
func (h *Handler) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, repository.NewHealthResponse(http.StatusOK, "Available"))
}
