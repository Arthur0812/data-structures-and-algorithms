package stacks

type node struct {
	element int
	next    *node
}

// Stack represents a stack data structure
type Stack struct {
	top *node
}

// CONSTRUCTORS

// EmptyStack is an empty stack
var EmptyStack = &Stack{}

func push(element int, stack *Stack) *Stack {
	newNode := &node{element: element, next: stack.top}
	return &Stack{top: newNode}
}

// SELECTORS

func top(stack *Stack) int {
	return stack.top.element
}

func pop(stack *Stack) *Stack {
	return &Stack{top: stack.top.next}
}

// CONDITION

func isEmpty(stack *Stack) bool {
	return stack.top == nil
}
