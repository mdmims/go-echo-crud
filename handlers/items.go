package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"goTestApi/models"
	"goTestApi/repository"

	"github.com/labstack/echo/v4"
)

// getAllItems godoc
// @Summary fetch all items
// @Accept json
// @Produce json
// @Success 200 {object} []repository.itemsResponse
// @Failure 400 {object} repository.Error
// @Failure 404 {object} repository.Error
// @Failure 500 {object} repository.Error
// @Router /items [get]
func (h *Handler) getAllItems(c echo.Context) error {
	// retrieve data from DB
	items, err := h.itemStore.GetAll()

	// check for errors or empty result
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, repository.NotFound())
		}
		return echo.NewHTTPError(http.StatusUnprocessableEntity, repository.NewError(err, http.StatusUnprocessableEntity, repository.EINTERNAL))
	}

	return c.JSON(http.StatusOK, items)
}

// createItem godoc
// @Summary create single item
// @Accept json
// @Produce json
// @Param user body repository.itemsResponse true "Item details to create record."
// @Success 200 {object} repository.itemsResponse
// @Failure 400 {object} repository.Error
// @Failure 404 {object} repository.Error
// @Failure 500 {object} repository.Error
// @Router /items [post]
func (h *Handler) createItem(c echo.Context) error {
	var m models.Items

	if err := c.Bind(&m); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, repository.NewError(err, http.StatusUnprocessableEntity, repository.EMALFORMED))
	}

	i, err := h.itemStore.Create(&m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, repository.NewError(err, http.StatusUnprocessableEntity, repository.EINTERNAL))
	}

	return c.JSON(http.StatusOK, repository.NewitemsResponse(i))
}

// getItem godoc
// @Summary retrieve single item
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} repository.itemsResponse
// @Failure 400 {object} repository.Error
// @Failure 404 {object} repository.Error
// @Failure 500 {object} repository.Error
// @Router /item/{id} [get]
func (h *Handler) getItem(c echo.Context) error {
	// retrieve path param value and convert to int
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// query for results
	i, err := h.itemStore.GetById(idd)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, repository.NotFound())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, repository.NewError(err, http.StatusInternalServerError, repository.EINTERNAL))
	}

	return c.JSON(http.StatusOK, repository.NewitemsResponse(i))
}
