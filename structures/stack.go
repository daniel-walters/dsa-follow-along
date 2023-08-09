package structures

import (
	"errors"
)

type Stack[T any] struct {
	head   *stackNode[T]
	length int
}

type stackNode[T any] struct {
	value T
	prev  *stackNode[T]
}

func newStackNode[T any](item T) *stackNode[T] {
	return &stackNode[T]{value: item}
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

var errorStackEmpty = errors.New("Stack is empty")

func (s *Stack[T]) Push(item T) {
	s.length++
	node := newStackNode(item)

	if s.head == nil {
		s.head = node
        return
	}

	node.prev = s.head
	s.head = node
}

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		return *(new(T)), errorStackEmpty
	}

	s.length--

	head := s.head
	s.head = head.prev
	head.prev = nil

	return head.value, nil
}

func (s Stack[T]) Peek() (T, error) {
	if s.head == nil {
		return *(new(T)), errorStackEmpty
	}

	return s.head.value, nil
}

func (s Stack[T]) Length() int {
	return s.length
}
