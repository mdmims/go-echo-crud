package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	// items endpoints
	items := v1.Group("/items")
	items.GET("", h.getAllItems)
	items.GET("/:id", h.getItem)
	items.POST("", h.createItem)
}
