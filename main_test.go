package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(Handler())
	defer ts.Close()
	res, err := ts.Client().Get(ts.URL + "/?name=fardean")
	if err != nil {
		log.Fatalf("error sending request: %s\n", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, res.StatusCode)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error during reading body: %s\n", err.Error())

	}
	if string(body) != "hello fardean\n" {
		log.Fatalf("want body to be: %s", "hello fardean\n")
	}
}
