package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInfo(t *testing.T) {
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	// response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Info)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	exspected := `{"version":"v:0.0.1"}`
	if rr.Body.String() != exspected {
		t.Errorf("handler returned unexspected body: got %v want %v",
			rr.Body.String(), exspected)
	}

}
