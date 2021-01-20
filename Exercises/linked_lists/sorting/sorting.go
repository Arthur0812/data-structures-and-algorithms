package sorting

type node struct {
	data int
	next *node
}

// LinkedList represents a linked list data structure
type LinkedList struct {
	head   *node
	length int
}

// EmptyList is simply an empty list, with no head and length 0
var EmptyList = &LinkedList{}

// MakeList creates a new linked list by combining the given linked list and the given data
func MakeList(data int, list *LinkedList) *LinkedList {
	n := &node{data: data, next: list.head}
	return &LinkedList{head: n, length: list.length + 1}
}

// Head returns the data stored in the head of the given list
func head(list *LinkedList) int {
	return list.head.data
}

// Rest returns the given list without its head
func rest(list *LinkedList) *LinkedList {
	return &LinkedList{head: list.head.next, length: list.length - 1}
}

func length(list *LinkedList) int {
	return list.length
}

// IsEmpty reports whether the given list is empty
func isEmpty(list *LinkedList) bool {
	return list.head == nil && list.length == 0
}

// Append appends the the new list to the first fiven list
func Append(list, new *LinkedList) *LinkedList {
	if isEmpty(list) {
		return new
	}
	return MakeList(head(list), Append(rest(list), new))
}

func MergeSort(list *LinkedList) *LinkedList {
	if isEmpty(list) || isEmpty(rest(list)) {
		return list
	}
	middle := getMiddle(list)
	next := rest(middle)

	middle.head.next = nil
	return list
}

func getMiddle(list *LinkedList) *LinkedList {
	curr, ln := list.head, length(list)
	for i := 0; i < ln && curr != nil; curr, i = curr.next, i+1 {

	}
	return &LinkedList{head: curr}
}
