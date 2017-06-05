package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplyCORS(t *testing.T) {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {})
	testserver := httptest.NewServer(applyCORS(http.DefaultServeMux))

	req, _ := http.NewRequest("OPTIONS", testserver.URL+"/", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("got %#v, want %#v", got, want)
	}

	if v, ok := resp.Header["Access-Control-Allow-Origin"]; !ok || v[0] != "*" {
		t.Fatalf("expected origin CORS header: got %s", resp.Header)
	}
	if v, ok := resp.Header["Access-Control-Allow-Methods"]; !ok || v[0] != "GET,HEAD,OPTIONS" {
		t.Fatalf("expected origin CORS header: got %s", resp.Header)
	}
}
