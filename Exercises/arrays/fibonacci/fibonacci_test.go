package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21}

	got := Fibonacci(8)

	if l1, l2 := len(got), len(expected); l1 != l2 {
		t.Fatalf("fibonacci sequence: incorrect length: expected: %v, got: %v", l2, l1)
	}

	for i, e := range got {
		if expected[i] != e {
			t.Fatalf("fibonacci sequence: incorrect number: expected: %v, got: %v", expected[i], e)
		}
	}
}
