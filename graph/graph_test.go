package graph_test

import (
	"dsa/graph"
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	graph_matrix := graph.WeightedAdjacencyMatrix{
		{0, 1, 4, 5, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 2, 0},
		{0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0},
	}

	expected := []int{0, 3, 4}
	actual, err := graph.BFS(graph_matrix, 0, 4)

	if err != nil {
		t.Error("Expected to find a path from 0 to 4")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}

    if actual, err = graph.BFS(graph_matrix, 3, 0); err == nil {
        t.Errorf("Expected error, receieved %+v", actual)
    }
}
