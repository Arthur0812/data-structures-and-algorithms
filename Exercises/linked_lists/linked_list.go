package linked_lists

type node struct {
	data int
	next *node
}

// LinkedList represents a linked list data structure
type LinkedList struct {
	Head   *node
	length int
}

// EmptyList is simply an empty list, with no head and length 0
var EmptyList = &LinkedList{}

// MakeList creates a new linked list by combining the given linked list and the given data
func MakeList(data int, list *LinkedList) *LinkedList {
	n := &node{data: data, next: list.Head}
	return &LinkedList{Head: n, length: list.length + 1}
}

// Head returns the data stored in the head of the given list
func Head(list *LinkedList) int {
	return list.Head.data
}

// Rest returns the given list without its head
func Rest(list *LinkedList) *LinkedList {
	return &LinkedList{Head: list.Head.next, length: list.length - 1}
}

// Length returns the given list's length of nodes
func Length(list *LinkedList) int {
	return list.length
}

// IsEmpty reports whether the given list is empty
func IsEmpty(list *LinkedList) bool {
	return list.Head == nil && list.length == 0
}

// Append appends the the new list to the first fiven list
func Append(list, new *LinkedList) *LinkedList {
	if IsEmpty(list) {
		return new
	}
	return MakeList(Head(list), Append(Rest(list), new))
}
