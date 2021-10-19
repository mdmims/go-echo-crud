package handlers

import (
	"database/sql"
	"net/http"

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
			return c.JSON(http.StatusNotFound, repository.Errorf("404", "No items found."))
		}
		return c.JSON(http.StatusInternalServerError, repository.Errorf("500", "Internal Server Error."))
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
		return c.JSON(http.StatusUnprocessableEntity, repository.Errorf("422", "Schema validation error."))
	}

	i, err := h.itemStore.Create(&m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, repository.NewitemsResponse(i))
}