package main

import (
	"github.com/hassanzreik/travel-path/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8081"))
}
