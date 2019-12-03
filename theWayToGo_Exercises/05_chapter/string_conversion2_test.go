package main

import (
	"testing"
)

func TestCanConvertString(t *testing.T) {

	result1, err1 := Converter("1234")

	if result1 != 1234 {
		t.Log("Failed to concert string")
		t.Fail()
	}

	if err1 != nil {
		t.Error("Failed to concert string")
	}

	result2, _ := Converter("1.12")

	if result2 != 0 {
		t.Error("Failed to return 0 for invalid value")
	}

}