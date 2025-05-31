package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestTravelPathHandler(t *testing.T) {
	e := echo.New()
	tickets := [][]string{
		{"LAX", "DXB"},
		{"JFK", "LAX"},
		{"SFO", "SJC"},
		{"DXB", "SFO"},
	}
	expected := []string{"JFK", "LAX", "DXB", "SFO", "SJC"}

	reqBody, _ := json.Marshal(tickets)
	req := httptest.NewRequest(http.MethodPost, "/travel-path", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := TravelPathHandler(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", rec.Code)
	}

	var actual []string
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if len(actual) != len(expected) {
		t.Fatalf("Length mismatch: expected %v, got %v", expected, actual)
	}

	if !slices.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestTravelPathHandler_EmptyInput(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/travel-path", bytes.NewBuffer([]byte(`[]`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := TravelPathHandler(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	var result []string
	json.Unmarshal(rec.Body.Bytes(), &result)

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
}
