package doublylinkedlists

type node struct {
	element     int
	left, right *node
}

// DoublyLinkedList represents a doubly linked list data structure
type DoublyLinkedList struct {
	first, last *node
}

// CONSTRUCTORS

// EmptyList is simply an empty doubly linked list
var EmptyList = &DoublyLinkedList{}

// MakeListLeft takes an element and a doubly linked list and returns a new list with the given element added to the left of the original list
func MakeListLeft(element int, list *DoublyLinkedList) *DoublyLinkedList {
	newNode := &node{element: element, right: list.first}
	newList := &DoublyLinkedList{first: newNode, last: list.last}
	if !isEmpty(list) {
		newNode.right.left = newNode
	} else {
		newList.last = newNode
	}
	return newList
}

// MakeListRight takes an element and a doubly linked list and returns a new list with the given element added to the right of the original list
func MakeListRight(element int, list *DoublyLinkedList) *DoublyLinkedList {
	newNode := &node{element: element, left: list.last}
	newList := &DoublyLinkedList{first: list.first, last: newNode}
	if !isEmpty(list) {
		newNode.left.right = newNode
	} else {
		newList.first = newNode
	}
	return newList
}

// SELECTORS

func firstLeft(list *DoublyLinkedList) int {
	return list.first.element
}

func firstRight(list *DoublyLinkedList) int {
	return list.last.element
}

func restLeft(list *DoublyLinkedList) *DoublyLinkedList {
	newList := &DoublyLinkedList{first: list.first.right, last: list.last}
	if !isEmpty(newList) {
		newList.first.left = nil
	}
	return newList
}

func restRight(list *DoublyLinkedList) *DoublyLinkedList {
	newList := &DoublyLinkedList{first: list.first, last: list.last.left}
	if !isEmpty(newList) {
		newList.last.right = nil
	}
	return newList
}

// CONDITIONS

func isEmpty(list *DoublyLinkedList) bool {
	return list.first == nil || list.last == nil
}

var myList = MakeListLeft(12, MakeListLeft(5, MakeListLeft(5, MakeListRight(20, MakeListRight(18, MakeListRight(7, EmptyList))))))
