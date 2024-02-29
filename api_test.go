package main_test

import (
	"encoding/json"
	main "github.com/xerardoo/poc-golang"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllTasks(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.GetAllTasks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response []main.Task
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error parsing response body: %v", err)
	}

	expectedLength := 2
	if len(response) != expectedLength {
		t.Errorf("handler returned unexpected task count: got %v want %v",
			len(response), expectedLength)
	}
}
