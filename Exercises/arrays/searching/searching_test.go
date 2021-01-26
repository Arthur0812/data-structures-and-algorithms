package searching

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	expected := 4
	testArr, x := []int{3, 10, 23, 18, 84, 30, 31, 53, 99, 1}, 84

	if got := LinearSearch(testArr, x); got != expected {
		t.Fatalf("linear search: incorrect index: expected: %v, got: %v", expected, got)
	}
}

func TestBinarySearch(t *testing.T) {
	expected := 6
	sortedArr, x := []int{1, 3, 14, 23, 58, 74, 84, 89, 92, 101, 110}, 84

	if got := BinarySearch(sortedArr, x); got != expected {
		t.Fatalf("linear search: incorrect index: expected: %v, got: %v", expected, got)
	}
}
