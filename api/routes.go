package api

import (
	"github.com/hassanzreik/travel-path/api/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/travel-path", handlers.TravelPathHandler)
}
