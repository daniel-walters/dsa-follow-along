package search_test

import (
	"dsa/search"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	array := []int{1, 3, 6, 9, 11}
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
