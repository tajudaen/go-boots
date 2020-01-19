package domain

import (
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("we were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("we were expecting an error when user id is 0")
	}
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	if err != nil {
		t.Error("we were not expecting an error")
	}

	if user == nil {
		t.Error("we were not expecting a user with id 123 to be nil")
	}
}