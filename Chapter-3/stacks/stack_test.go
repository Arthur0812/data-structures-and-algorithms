package stacks

import "testing"

func TestStacks(t *testing.T) {
	expected := []int{5, 2, 4, 1, 3}

	example := []int{3, 1, 4, 2, 5}

	var myStack = EmptyStack
	for _, n := range example {
		myStack = push(n, myStack)
	}

	for i := 0; !isEmpty(myStack) && i < len(expected); i++ {
		if top(myStack) != expected[i] {
			t.Fatalf("stacks: incorrect element: expected: %v, got: %v", expected[i], top(myStack))
		}
		myStack = pop(myStack)
	}

	testStack := push(3, push(9, EmptyStack))
	if !isEmpty(EmptyStack) {
		t.Fatalf("stacks: empty stack should be empty")
	}
	if isEmpty(testStack) {
		t.Fatalf("stacks: stack should not be empty")
	}
	if top(testStack) != 3 {
		t.Fatalf("stacks: top element does not match")
	}
	if pop(push(10, testStack)).top != testStack.top {
		t.Fatalf("stacks: popped stack does not match")
	}
}
