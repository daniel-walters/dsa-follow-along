package structures_test

import (
	"dsa/structures"
	"testing"
)

func TestLinkedLinked(t *testing.T) {
	list := structures.NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Prepend(3)

	if list.Length() != 3 {
		t.Errorf("Incorrect length, recieved: %d", list.Length())
	}

	if head, _ := list.Get(0); head != 3 {
		t.Errorf("Incorrect head, receieved: %d", head)
	}
	if tail, _ := list.Get(2); tail != 2 {
		t.Errorf("Incorrect tail, receieved: %d", tail)
	}

	value, err := list.RemoveAt(1)
	if err != nil {
		t.Errorf("Got error: %s", err)
	}
	if value != 1 {
		t.Errorf("Incorrect value, receieved: %v", value)
	}

	list.Append(4)
	list.Append(5)

	err = list.InsertAt(6, 3)
	if err != nil {
		t.Errorf("got error :%s", err)
	}
	value, err = list.Get(3)
	if err != nil {
		t.Errorf("got error :%s", err)
	}
	if value != 6 {
		t.Errorf("incorrect value, recieved: %d", value)
	}
}

func TestQueue(t *testing.T) {
	queue := structures.NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if queue.Length() != 3 {
		t.Errorf("Incorrect length, recieved %d", queue.Length())
	}

	if v, _ := queue.Peek(); v != 1 {
		t.Errorf("Incorrect value, recieved %d", v)
	}

	if v, _ := queue.Dequeue(); v != 1 {
		t.Errorf("Incorrect value, recieved %d", v)
	}

	if v, _ := queue.Peek(); v != 2 {
		t.Errorf("Incorrect value, recieved %d", v)
	}

	if queue.Length() != 2 {
		t.Errorf("Incorrect length, recieved %d", queue.Length())
	}

	queue.Dequeue()
	queue.Dequeue()

	if _, e := queue.Dequeue(); e == nil {
		t.Error("Expected error, recieved nil")
	}
}
	}

}
