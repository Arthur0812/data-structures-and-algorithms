package binarysearch

import (
	"sort"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	example := []int{3, 14, 16, 92, 62, 54, 1, 78, 4, 20, 17, 40, 6}

	// IMPORTANT: SLICE HAS TO BE ORDERED, in order for the algorithm to work expectedly
	sort.Ints(example)
	// ordered =[1, 3, 4, 6, 14, 16, 17, 20, 40, 54, 62, 78, 92]

	x := 55
	expected := -1
	if got := binarySearch(example, x); got != expected {
		t.Fatalf("linear search: incorrect index: expected: %v, got: %v", expected, got)
	}

	y := 54
	expected = 9
	if got := binarySearch(example, y); got != expected {
		t.Fatalf("linear search: incorrect index: expected: %v, got: %v", expected, got)
	}
}
