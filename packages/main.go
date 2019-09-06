package main

import (
	"fmt"
	"strings"
	"github.com/tajud99n/web_dev/packages/math"
)

func main() {
	//*****STRINGS**********

	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test", "es")) // true

	// func Count(s, sep string) int
	fmt.Println(strings.Count("test", "t")) // 2

	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "te")) // true

	// func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("test", "st")) // true

	// func Index(s, sep string) int
	fmt.Println(strings.Index("test", "e")) // 1

	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"a","b","c"}, "-")) // "a-b-c"

	// func Repeat(s string, count int) string
	fmt.Println(strings.Repeat("A", 3)) // "AAA"

	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("aaaa", "a", "b", 2)) // "bbaa"

	// func Split(s, sep string) []string
	fmt.Println(strings.Split("a-b-c-d-e", "-")) // []string{"a","b","c","d","e"}

	// func ToLower(s string) string
	fmt.Println(strings.ToLower("TEST")) // "test"

	// func ToUpper(s string) string
	fmt.Println(strings.ToUpper("test")) // "TEST"

	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg) // 2.5
}