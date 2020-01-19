package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	els := []int{9, 7, 8, 5, 6}

	els = BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
}

func getElement(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}

	return result
}

func BenchmarkBubbleSort(b *testing.B) {
	els := getElement(100)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
