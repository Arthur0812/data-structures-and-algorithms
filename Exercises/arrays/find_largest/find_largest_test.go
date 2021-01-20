package findlargest

import "testing"

func TestFindLargest(t *testing.T) {
	expected := 6
	test := []int{20, 5, 10, 23, 34, 100, 260, 12, 34}

	if got := FindLargest(test); got != expected {
		t.Fatalf("find largest: incorrect index: expected: %v, got: %v", expected, got)
	}
}
