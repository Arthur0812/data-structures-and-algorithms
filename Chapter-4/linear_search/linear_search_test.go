package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	example := []int{3, 14, 16, 92, 62, 54, 1, 78, 4, 20, 17, 40, 6}

	x := 55
	expected := -1
	if got := linearSearch(example, x); got != expected {
		t.Fatalf("linear search: incorrect index: expected: %v, got: %v", expected, got)
	}

	y := 54
	expected = 5
	if got := linearSearch(example, y); got != expected {
		t.Fatalf("linear search: incorrect index: expected: %v, got: %v", expected, got)
	}
}
