package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"rssfeed/controllers"
	"testing"
)

func TestSearchRss(t *testing.T) {

	var jsonStr = []byte(`{"topic": "virus"}`)

	req, err := http.NewRequest("POST", "/search", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.SearchRssFeed)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSearchRssNoTopic(t *testing.T) {

	var jsonStr = []byte(`{"topic": ""}`)

	req, err := http.NewRequest("POST", "/search", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.SearchRssFeed)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
