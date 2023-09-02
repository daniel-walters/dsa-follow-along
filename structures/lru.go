package structures

import (
	"errors"
)

type lruNode[T any] struct {
	value T
	next  *lruNode[T]
	prev  *lruNode[T]
}

func newLRUNode[T any](value T) *lruNode[T] {
	return &lruNode[T]{value: value}
}

type LRU[K comparable, V any] struct {
	head          *lruNode[V]
	tail          *lruNode[V]
	length        int
	lookup        map[K]*lruNode[V]
	reverseLookup map[*lruNode[V]]K
	capacity      int
}

func NewLRU[K comparable, V any](capacity int) LRU[K, V] {
	if capacity == 0 {
		capacity = 10
	}

	return LRU[K, V]{
		head:          nil,
		tail:          nil,
		length:        0,
		lookup:        map[K]*lruNode[V]{},
		reverseLookup: map[*lruNode[V]]K{},
		capacity:      capacity,
	}
}

func (lru *LRU[K, V]) Update(key K, value V) {
	// does it exist?
	node, nodeExists := lru.lookup[key]

	// if not need to insert
	if !nodeExists {
		node = newLRUNode(value)
		lru.length += 1
		lru.prepend(node)
		lru.trimCache()

		lru.lookup[key] = node
		lru.reverseLookup[node] = key
	} else {
		lru.detach(node)
		lru.prepend(node)
		node.value = value
	}
}

func (lru *LRU[K, V]) Get(key K) (V, error) {
	node, nodeExists := lru.lookup[key]
	if !nodeExists {
		return *new(V), errors.New("No node associated to key")
	}

	lru.detach(node)
	lru.prepend(node)

	return node.value, nil
}

func (lru *LRU[K, V]) trimCache() {
	if lru.length <= lru.capacity {
		return
	}

	tail := lru.tail
	lru.detach(lru.tail)

	key := lru.reverseLookup[tail]
	delete(lru.lookup, key)
	delete(lru.reverseLookup, tail)

	lru.length -= 1
}

func (lru *LRU[K, V]) detach(node *lruNode[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	// if lru.length == 1 {
	// 	lru.tail = nil
	// 	lru.head = nil
	// }

	if lru.head == node {
		lru.head = lru.head.next
	}

	if lru.tail == node {
		lru.tail = lru.tail.prev
	}

	node.next = nil
	node.prev = nil

}

func (lru *LRU[K, V]) prepend(node *lruNode[V]) {
	if lru.head == nil {
		lru.head = node
		lru.tail = node

		return
	}

	node.next = lru.head
	lru.head.prev = node
	lru.head = node
}
