package main

import (
	"reflect"
	"testing"
)

func TestNewLot(t *testing.T) {
	expectedRegularSpaces := 3
	expectedADASpaces := 4
	expectedTotalSpaces := expectedADASpaces + expectedRegularSpaces
	l, r := newLot(expectedRegularSpaces, expectedADASpaces)

	if len(l) != expectedTotalSpaces {
		t.Errorf("Expected new lot of %v spaces.  but got %v", expectedTotalSpaces, len(l))
	}

	if reflect.TypeOf(r) != reflect.TypeOf("string") {
		t.Errorf("Expected a response of string")
	}

}
