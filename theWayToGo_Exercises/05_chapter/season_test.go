package main

import (
	"testing"
)

func TestCanReturnSeason(t *testing.T) {
	result := Season(0)

	if result != "season1" {
		t.Error("Failed to ...")
	}
}