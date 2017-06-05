package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/platelk/fizzbuzz-api/core"
)

func TestVersionRoute(t *testing.T) {
	httpService := CreateFizzBuzzService()
	testserver := httptest.NewServer(httpService.GetHandler())

	req, _ := http.NewRequest("GET", testserver.URL + "/fizzbuzz/v1/version", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	var versionMsg VersionMessage
	err = json.Unmarshal(data, &versionMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	if got, want := versionMsg.Version, httpService.Version(); got != want {
		t.Fatalf("Wrong version return, got %s but want %s", got, want)
	}
}

func TestFizzBuzzRoute(t *testing.T) {
	httpService := CreateFizzBuzzService()
	testserver := httptest.NewServer(httpService.GetHandler())

	from, to, mul1, mul2, s1, s2 := 1, 15, 3, 5, "fizz", "buzz"

	req, _ := http.NewRequest("GET",
		testserver.URL + fmt.Sprintf("/fizzbuzz/v1/fizzbuzz?to=%d&mul1=%d&mul2=%d&word1=%s&word2=%s", to, mul1, mul2, s1, s2),
		nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	var fizzBuzzMsg FizzBuzzMessage
	err = json.Unmarshal(data, &fizzBuzzMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	want, _ := core.FizzBuzz(from, to, mul1, mul2, s1, s2)
	if got := fizzBuzzMsg.Response; !reflect.DeepEqual(got, want) {
		t.Fatalf("Wrong version return, got %s but want %s", got, want)
	}
}