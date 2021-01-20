package linkedlists

type node struct {
	element int
	next    *node
}

// List represents a linked list "wrapper"
type List struct {
	first *node
}

// CONSTRUCTORS

// EmptyList is simply an empty linked list
var EmptyList = &List{}

// MakeList puts an element on top of an existing linked list and returns the existing list
func MakeList(element int, list *List) *List {
	newNode := &node{element: element, next: list.first}
	return &List{first: newNode}
}

// SELECTORS

func first(list *List) int {
	return list.first.element
}

func rest(list *List) *List {
	return &List{first: list.first.next}
}

// CONDITIONS

func isEmpty(list *List) bool {
	return list.first == nil
}

var myList = MakeList(3, MakeList(1, MakeList(4, MakeList(2, MakeList(5, EmptyList)))))

// MUTATORS

func replaceFirst(newElement int, list *List) {
	if isEmpty(list) {
		return
	}
	list.first.element = newElement
}

func replaceRest(newList *List, list *List) {
	if isEmpty(list) {
		return
	}
	list.first.next = newList.first
}

// RECURIVE ALGORITHMS

func last(list *List) (int, bool) {
	if isEmpty(list) {
		return 0, false
	}
	if isEmpty(rest(list)) {
		return first(list), true
	}
	return last(rest(list))
}

func append(list *List, new *List) *List {
	if isEmpty(list) {
		return new
	}
	return MakeList(first(list), append(rest(list), new))
}
