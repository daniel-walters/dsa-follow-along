package structures

import "errors"

type linkedListNode[T any] struct {
	value T
	next  *linkedListNode[T]
	prev  *linkedListNode[T]
}

type LinkedList[T any] struct {
	head   *linkedListNode[T]
	tail   *linkedListNode[T]
	length int
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{length: 0}
}

func newListNode[T any](value T) *linkedListNode[T] {
	return &linkedListNode[T]{value: value}
}

func (list *LinkedList[T]) InsertAt(item T, index int) error {
	currentNode, err := list.traverse(index)
	if err != nil {
		return err
	}

	newNode := newListNode(item)
	list.length++

    prevNode := currentNode.prev
    newNode.prev = prevNode
    newNode.next = currentNode

    prevNode.next = newNode
    currentNode.prev = newNode

	return nil
}

func (list LinkedList[T]) Length() int {
	return list.length
}

func (list *LinkedList[T]) RemoveAt(index int) (T, error) {
	node, err := list.traverse(index)
	if err != nil {
		return *(new(T)), err
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil

	list.length--

	return node.value, nil
}

func (list *LinkedList[T]) Append(item T) {
	node := newListNode(item)
	list.length++

	if list.head == nil && list.tail == nil {
		list.head = node
		list.tail = node

		return
	}

	node.prev = list.tail
	list.tail.next = node
	list.tail = node
}

func (list *LinkedList[T]) Prepend(item T) {
	node := newListNode(item)
	list.length++

	if list.head == nil && list.tail == nil {
		list.head = node
		list.tail = node

		return
	}

	node.next = list.head
	list.head.prev = node
	list.head = node
}

func (list LinkedList[T]) Get(index int) (T, error) {
	node, err := list.traverse(index)
	if err != nil {
		return *(new(T)), err
	}

	return node.value, nil

}

func (list LinkedList[T]) traverse(index int) (*linkedListNode[T], error) {
	var node *linkedListNode[T] = list.head

	for i := 1; i <= index; i++ {
		if node == nil {
			return nil, errors.New("Index out of bounds")
		}
		node = node.next
	}

	return node, nil
}
