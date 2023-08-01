package search_test

import (
	"dsa/search"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	array := []int{11, 3, 9, 6, 1}
	expected := true
	actual := search.LinearSearch(array, 6)
	if actual != expected {
		t.Errorf("Expected: %t, Received: %t", expected, actual)
	}

	expected = false
	actual = search.LinearSearch(array, 100)
	if actual != expected {
		t.Errorf("Expected: %t, Received: %t", expected, actual)
	}
}

func TestBinarySearch(t *testing.T) {
	array := []int{1, 3, 6, 9, 11, 17, 21, 35, 57}
	expected := true
	actual := search.BinarySearch(array, 9)
	if actual != expected {
		t.Errorf("Expected: %t, Received: %t", expected, actual)
	}
	actual = search.BinarySearch(array, 35)
	if actual != expected {
		t.Errorf("Expected: %t, Received: %t", expected, actual)
	}

	expected = false
	actual = search.BinarySearch(array, 100)
	if actual != expected {
		t.Errorf("Expected: %t, Received: %t", expected, actual)
	}
}
