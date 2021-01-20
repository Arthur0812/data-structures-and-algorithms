package linkedlists

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	expected := []int{3, 1, 4, 2, 5}

	for i := 0; !isEmpty(myList) && i < len(expected); i++ {
		if first(myList) != expected[i] {
			t.Fatalf("linked lists: incorrect element: expected: %v, got: %v", expected[i], first(myList))
		}
		myList = rest(myList)
	}

	testList := MakeList(12, MakeList(8, MakeList(7, EmptyList)))

	if !isEmpty(EmptyList) {
		t.Fatal("linked lists: empty list should be empty")
	}
	if isEmpty(testList) {
		t.Fatal("linked lists: list should not be empty")
	}
	if first(testList) != 12 {
		t.Fatal("linked lists: first element does not match")
	}
	if rest(MakeList(100, testList)).first != testList.first {
		t.Fatal("linked lists: rest list does not match")
	}
}
