package main

import (
	"testing"
)

func TestCanReturnAPyramidStructure(t *testing.T) {
	result := Repeat("A", 3)

	if len(result) != 9 {
		t.Error("Failed to pyramid structure")
	}
}
