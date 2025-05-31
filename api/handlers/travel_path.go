package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	emptyInputErrMsg                 = "input is empty"
	invalidTicketErrMsg              = "invalid ticket format, must be a pair of strings"
	emptyDepartureDestinationErrMsg  = "departure and destination must not be empty"
	departureSameAsDestinationErrMsg = "departure cannot be the same as destination: "
	duplicateDepartureErrMsg         = "duplicate departure found "
	noValidStartingPointErrMsg       = "no valid starting point found"
	cycleDetectedErrMsg              = "cycle detected in path"
	incompleteItineraryErrMsg        = "incomplete itinerary path"
)

type flightTicketInput [][]string

func TravelPathHandler(c echo.Context) error {
	var tickets flightTicketInput
	if err := c.Bind(&tickets); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	path, err := buildTravelPath(tickets)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, path)
}

func buildTravelPath(flightTickets [][]string) ([]string, error) {
	if len(flightTickets) == 0 {
		return nil, errors.New(emptyInputErrMsg)
	}

	routeMap := make(map[string]string)
	hasIncoming := make(map[string]bool)

	for _, flightTicket := range flightTickets {
		if len(flightTicket) != 2 {
			return nil, errors.New(invalidTicketErrMsg)
		}
		dep, dest := flightTicket[0], flightTicket[1]
		if dep == "" || dest == "" {
			return nil, errors.New(emptyDepartureDestinationErrMsg)
		}
		if dep == dest {
			return nil, errors.New(departureSameAsDestinationErrMsg + dep)
		}
		if _, exists := routeMap[dep]; exists {
			return nil, errors.New(duplicateDepartureErrMsg + dep)
		}
		routeMap[dep] = dest
		hasIncoming[dest] = true
	}

	var start string
	for _, flightTicket := range flightTickets {
		if !hasIncoming[flightTicket[0]] {
			start = flightTicket[0]
			break
		}
	}
	if start == "" {
		return nil, errors.New(noValidStartingPointErrMsg)
	}

	var path []string
	visited := make(map[string]bool)

	for current := start; current != ""; current = routeMap[current] {
		if visited[current] {
			return nil, errors.New(cycleDetectedErrMsg)
		}
		visited[current] = true
		path = append(path, current)

		if _, ok := routeMap[current]; !ok {
			break
		}
	}

	if len(path) != len(flightTickets)+1 {
		return nil, errors.New(incompleteItineraryErrMsg)
	}

	return path, nil
}
