package sample

import (
	"golang-rest-boilerplate/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Create, middleware.Authentication)
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetById, middleware.Authentication)
	g.PUT("/:code", h.UpdateByCode, middleware.Authentication)
	g.DELETE("/:code", h.DeleteByCode, middleware.Authentication)
	// for endpoint doesn't need authentication, just delete the section middlewarre.Authentication like g.GET("/:id", h.GetById)
}
