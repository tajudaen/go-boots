package main

import (
	"math"
	"errors"
)

func Square(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("only positive float64 type is allowed")
	}
	s := math.Sqrt(n)
	return s, nil
}