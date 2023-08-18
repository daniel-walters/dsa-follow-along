package sort_test

import (
	"dsa/sort"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{11, 3, 9, 6, 1}
	expected := []int{1, 3, 6, 9, 11}
	sort.BubbleSort(array)
	if !reflect.DeepEqual(array, expected) {
		t.Errorf("Expected: %+v, Receieved: %+v", expected, array)
	}
}

func TestQuickSort(t *testing.T) {
	array := []int{11, 3, 9, 6, 1, 5, 9, 23, 2}
	expected := []int{1, 2, 3, 5, 6, 9, 9, 11, 23}
	sort.QuickSort(array)
	if !reflect.DeepEqual(array, expected) {
		t.Errorf("Expected: %+v, Receieved: %+v", expected, array)
	}
}
