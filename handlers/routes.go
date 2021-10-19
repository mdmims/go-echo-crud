package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	// ticket header
	ticket := v1.Group("/items")
	ticket.GET("", h.getAllItems)
	ticket.POST("", h.createItem)
}
