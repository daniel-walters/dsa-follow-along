package structures

import (
	"errors"
)

type queueNode[T any] struct {
	value T
	next  *queueNode[T]
}

func newQueueNode[T any](item T) *queueNode[T] {
	return &queueNode[T]{value: item}
}

type Queue[T any] struct {
	head   *queueNode[T]
	tail   *queueNode[T]
	length int
}

var ErrorEmptyQueue = errors.New("Queue is empty")

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	q.length++
	node := newQueueNode(item)

	if q.head == nil && q.tail == nil {
		q.head = node
		q.tail = node

		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.length == 0 {
		return *(new(T)), ErrorEmptyQueue
	}

	q.length--

	node := q.head
	q.head = node.next
	node.next = nil

	if q.length == 0 {
		q.tail = nil
	}

	return node.value, nil
}

func (q Queue[T]) Peek() (T, error) {
	if q.length == 0 {
		return *(new(T)), ErrorEmptyQueue
	}

	return q.head.value, nil
}

func (q Queue[T]) Length() int {
	return q.length
}
