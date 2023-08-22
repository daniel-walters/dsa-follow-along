package structures

import (
	"cmp"
	"errors"
)

type MinHeap[T cmp.Ordered] struct {
	data   []T
	length int
}

func NewMinHeap[T cmp.Ordered]() MinHeap[T] {
	return MinHeap[T]{
		data:   []T{},
		length: 0,
	}
}

func (heap MinHeap[T]) Length() int {
	return heap.length
}

func (heap *MinHeap[T]) Insert(item T) {
	heap.data = append(heap.data, item)
	heap.heapifyUp(heap.length)
	heap.length++
}

func (heap *MinHeap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parentIdx := parent(idx)
	parentVal := heap.data[parentIdx]
	curVal := heap.data[idx]

	if parentVal > curVal {
		heap.data[parentIdx] = curVal
		heap.data[idx] = parentVal

		heap.heapifyUp(parentIdx)
	}
}

func (heap *MinHeap[T]) heapifyDown(idx int) {
	leftIdx := leftChild(idx)
	rightIdx := rightChild(idx)

	if idx >= heap.length || leftIdx >= heap.length {
		return
	}

	curValue := heap.data[idx]
	leftValue := heap.data[leftIdx]
	if rightIdx >= heap.length && leftValue < curValue {
		heap.data[idx] = leftValue
		heap.data[leftIdx] = curValue

		heap.heapifyDown(leftIdx)
	} else if rightValue := heap.data[rightIdx]; leftValue >= rightValue {
		heap.data[rightIdx] = curValue
		heap.data[idx] = rightValue

		heap.heapifyDown(rightIdx)
	} else if rightValue > leftValue {
		heap.data[leftIdx] = curValue
		heap.data[idx] = leftValue

		heap.heapifyDown(leftIdx)
	}

}

func (heap *MinHeap[T]) Delete() (T, error) {
	if heap.length == 0 {
		return *new(T), errors.New("Heap is empty cannot delete")
	}

	out := heap.data[0]

	if heap.length == 1 {
		heap.data = []T{}
		heap.length = 0

		return out, nil
	}

    heap.length--
	heap.data[0] = heap.data[heap.length]
	heap.heapifyDown(0)

	return out, nil
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return 2*idx + 1
}

func rightChild(idx int) int {
	return 2*idx + 2
}
