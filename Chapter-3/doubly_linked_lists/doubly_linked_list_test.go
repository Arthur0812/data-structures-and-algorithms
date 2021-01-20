package doublylinkedlists

import "testing"

func TestDoublyLinkedList(t *testing.T) {
	expected := []int{12, 5, 5, 7, 18, 20}

	for i := len(expected) - 1; !isEmpty(myList) && i >= 0; i-- {
		if firstRight(myList) != expected[i] {
			t.Fatalf("doubly linked lists: incorrect element: expected: %v, got: %v", expected[i], firstRight(myList))
		}
		myList = restRight(myList)
	}

	testList1 := MakeListLeft(4, MakeListLeft(2, MakeListLeft(5, EmptyList)))
	testList2 := MakeListRight(11, MakeListRight(10, EmptyList))

	if !isEmpty(EmptyList) {
		t.Fatal("doubly linked lists: empty list should be empty")
	}
	if isEmpty(testList1) {
		t.Fatal("doubly linked lists: list 1 should not be empty")
	}
	if isEmpty(testList2) {
		t.Fatal("doubly linked lists: list 2 should not be empty")
	}

	if firstLeft(testList1) != 4 {
		t.Fatal("doubly linked lists: first element does not match")
	}
	if firstRight(testList2) != 11 {
		t.Fatal("doubly linked lists: last element does not match")
	}

	if restLeft(MakeListLeft(100, testList1)).first != testList1.first {
		t.Fatal("doubly linked lists: first rest does not match")
	}
	if restRight(MakeListRight(200, testList1)).first != testList1.first {
		t.Fatal("doubly linked lists: last rest does not match")
	}
}
