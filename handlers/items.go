package handlers

import (
    "database/sql"
    "net/http"

    "goTestApi/repository"

    "github.com/labstack/echo/v4"
)

// getTicket godoc
// @Summary GetFacts Ticket Header data
// @Description GetFacts Ticket Header
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Param ticketNumber path int true "ticketNumber"
// @Success 200 {object} repository.ticketResponse
// @Failure 400 {object} repository.Error
// @Failure 404 {object} repository.Error
// @Failure 500 {object} repository.Error
// @Router /tickets/{ticketNumber} [get]
func (h *Handler) getAllItems(c echo.Context) error {
    // retrieve data from DB
    i, err := h.itemStore.GetAll()

    // check for errors or empty result
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, repository.Errorf("404", "Not items found."))
        }
        return c.JSON(http.StatusInternalServerError, repository.Errorf("500", "Internal Server Error."))
    }

    return c.JSON(http.StatusOK, repository.NewitemsResponse(i))
    }
