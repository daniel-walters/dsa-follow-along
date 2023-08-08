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

	for i := 0; i < list.Length(); i++ {
		v, _ := list.Get(i)
		println(v)
	}

}
