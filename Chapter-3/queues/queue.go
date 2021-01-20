package queues

type node struct {
	element int
	next    *node
}

// Queue represents a queue data structure
type Queue struct {
	top, end *node
}

// CONSTRUCTORS

// EmptyQueue is simply an empty queue
var EmptyQueue = &Queue{}

func push(element int, queue *Queue) *Queue {
	newNode := &node{element: element}
	newQueue := &Queue{top: queue.top, end: queue.end}
	if !isEmpty(queue) {
		newQueue.end.next = newNode
	}
	newQueue.end = newNode
	return newQueue
}

// SELECTORS

func top(queue *Queue) int {
	return queue.top.element
}

func pop(queue *Queue) *Queue {
	return &Queue{top: queue.top.next, end: queue.end}
}

// CONDITIONS

func isEmpty(queue *Queue) bool {
	return queue.top == nil
}
