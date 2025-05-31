package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type flightTicketInput [][]string

func TravelPathHandler(c echo.Context) error {
	var tickets flightTicketInput
	if err := c.Bind(&tickets); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	path := buildTravelPath(tickets)
	return c.JSON(http.StatusOK, path)
}

func buildTravelPath(flightTickets [][]string) []string {
	if len(flightTickets) > 0 {
		return []string{}
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

	return path
}
