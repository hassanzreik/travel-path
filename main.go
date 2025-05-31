package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/travel-path", travelPath)
	e.Logger.Fatal(e.Start(":8081"))
}

type flightTicketsInput [][]string

func travelPath(c echo.Context) error {
	var flightTickets flightTicketsInput
	if err := c.Bind(&flightTickets); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	routeMap := make(map[string]string)
	hasIncoming := make(map[string]bool)

	for _, flightTicket := range flightTickets {
		src, dst := flightTicket[0], flightTicket[1]
		routeMap[src] = dst
		hasIncoming[dst] = true
	}

	var start string
	for _, flightTicket := range flightTickets {
		if !hasIncoming[flightTicket[0]] {
			start = flightTicket[0]
			break
		}
	}

	var path []string
	for curr := start; curr != ""; curr = routeMap[curr] {
		path = append(path, curr)
		if _, ok := routeMap[curr]; !ok {
			break
		}
	}

	return c.JSON(http.StatusOK, path)
}
