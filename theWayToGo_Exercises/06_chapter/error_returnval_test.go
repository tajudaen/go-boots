package main

import (
	"testing"
)

type e error

func TestCanReturnError(t *testing.T) {
	result, _ := Square(16)

	if result != 4 {
		t.Error("Failed to return square root")
	}

	_, err := Square(-9)

	if err == nil {
		t.Error("Failed to return error")
	}
}
