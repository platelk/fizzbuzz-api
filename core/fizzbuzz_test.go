package core

import (
	"reflect"
	"testing"
)

func TestFizzBuzzSimpleOK(t *testing.T) {
	expected := []string{"1", "2", "fizz", "4", "buzz"}
	res, err := FizzBuzz(1, 5, 3, 5, "fizz", "buzz")
	if err != nil {
		t.Fatalf("Got an error on normal test: %s", err.Error())
	}
	if got := reflect.DeepEqual(expected, res); !got {
		t.Fatalf("%s is not equal to expected (%s)", res, expected)
	}
}

func TestFizzBuzzDoubleOk(t *testing.T) {
	expected := []string{"1", "fizz", "buzz", "fizz", "5", "fizzbuzz"}
	res, err := FizzBuzz(1, 6, 2, 3, "fizz", "buzz")
	if err != nil {
		t.Fatalf("Got an error on normal test: %s", err.Error())
	}
	if got := reflect.DeepEqual(expected, res); !got {
		t.Fatalf("%s is not equal to expected (%s)", res, expected)
	}
}


func TestFizzBuzzErrorFromGreaterThanTo(t *testing.T) {
	_, err := FizzBuzz(6, 5, 2, 3, "fizz", "buzz")
	if err == nil {
		t.Fatalf("Got no error when from is greater than to (%d > %d)", 6, 5)
	}
}

func TestFizzBuzzErrorMultipleEqualZero(t *testing.T) {
	_, err := FizzBuzz(1, 5, 0, 3, "fizz", "buzz")
	if err == nil {
		t.Fatalf("Got no error when one multiple is equal to zero")
	}
}
