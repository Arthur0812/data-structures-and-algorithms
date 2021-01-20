package sorting

import "testing"

func TestBubbleSort(t *testing.T) {
	expected := []int{3, 4, 4, 6, 10, 20, 24, 35, 78, 99}

	test := []int{78, 10, 20, 6, 4, 3, 24, 4, 99, 35}

	BubbleSort(test)

	if l1, l2 := len(expected), len(test); l1 != l2 {
		t.Fatalf("bubble sort: incorrect length: expected: %v, got: %v", l1, l2)
	}

	for i, el := range test {
		if expected[i] != el {
			t.Fatalf("bubble sort: incorrect length: expected: %v, got: %v", expected[i], el)
		}
	}
}

func TestSelectionSortInt(t *testing.T) {
	expected := []int{3, 4, 4, 8, 10, 20, 45}

	test := []int{4, 10, 3, 45, 8, 4, 20}

	SelectionSortInt(test)

	if l1, l2 := len(expected), len(test); l1 != l2 {
		t.Fatalf("selection sort: incorrect length: expected: %v, got: %v", l1, l2)
	}

	for i, el := range test {
		if el != expected[i] {
			t.Fatalf("selection sort: incorrect element: expected: %v, got: %v", expected[i], el)
		}
	}
}

func TestSelectionSortString(t *testing.T) {
	expected := []string{"apple", "banana", "lemon", "orange", "pineapple"}

	test := []string{"banana", "pineapple", "orange", "apple", "lemon"}

	SelectionSortString(test)

	if l1, l2 := len(expected), len(test); l1 != l2 {
		t.Fatalf("selection sort: incorrect length: expected: %v, got: %v", l1, l2)
	}

	for i, el := range test {
		if el != expected[i] {
			t.Fatalf("selection sort: incorrect element: expected: %v, got: %v", expected[i], el)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	expected := []int{3, 4, 4, 8, 10, 20, 45, 60}

	test := []int{4, 10, 3, 45, 8, 60, 4, 20}

	InsertionSort(test)

	if l1, l2 := len(expected), len(test); l1 != l2 {
		t.Fatalf("insertion sort: incorrect length: expected: %v, got: %v", l1, l2)
	}

	for i, el := range test {
		if el != expected[i] {
			t.Fatalf("insertion sort: incorrect element: expected: %v, got: %v", expected[i], el)
		}
	}
}

func TestMergeSort(t *testing.T) {
	expected := []int{3, 4, 4, 8, 10, 20, 45, 60}

	test := []int{4, 10, 3, 45, 8, 60, 4, 20}

	MergeSort(test)

	if l1, l2 := len(expected), len(test); l1 != l2 {
		t.Fatalf("merge sort: incorrect length: expected: %v, got: %v", l1, l2)
	}

	for i, el := range test {
		if el != expected[i] {
			t.Fatalf("merge sort: incorrect element: expected: %v, got: %v", expected[i], el)
		}
	}
}

func TestHeapSort(t *testing.T) {
	expected := []int{3, 4, 4, 8, 10, 20, 45, 60}

	test := []int{4, 10, 3, 45, 8, 60, 4, 20}

	HeapSort(test)

	if l1, l2 := len(expected), len(test); l1 != l2 {
		t.Fatalf("merge sort: incorrect length: expected: %v, got: %v", l1, l2)
	}

	for i, el := range test {
		if el != expected[i] {
			t.Fatalf("merge sort: incorrect element: expected: %v, got: %v", expected[i], el)
		}
	}
}
