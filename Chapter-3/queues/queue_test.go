package queues

import "testing"

func TestQueue(t *testing.T) {
	expected := []int{10, 13, 7, 4, 1}
	example := []int{10, 13, 7, 4, 1}

	var myQueue = EmptyQueue
	for _, n := range example {
		myQueue = push(n, myQueue)
	}

	for i := 0; !isEmpty(myQueue) && i < len(expected); i++ {
		if top(myQueue) != expected[i] {
			t.Fatalf("queues: incorrect element: expected: %v, got: %v", expected[i], top(myQueue))
		}
		myQueue = pop(myQueue)
	}
}
