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

func TestDFS(t *testing.T) {
	graph_list := graph.WeightedAdjacencyList{
		{
			graph.GraphEdge{To: 1, Weight: 1},
			graph.GraphEdge{To: 2, Weight: 4},
			graph.GraphEdge{To: 3, Weight: 5},
		},
		{graph.GraphEdge{To: 0, Weight: 1}},
		{graph.GraphEdge{To: 3, Weight: 2}},
		{graph.GraphEdge{To: 4, Weight: 5}},
		{},
	}

	expected := []int{0, 2, 3, 4}
	actual, err := graph.DFS(graph_list, 0, 4)

	if err != nil {
		t.Error("Expected to find a path from 0 to 4")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}

	if actual, err = graph.DFS(graph_list, 3, 0); err == nil {
		t.Errorf("Expected error, receieved %+v", actual)
	}
}

func TestDijstras(t *testing.T) {
	graph_list := graph.WeightedAdjacencyList{
		{
			graph.GraphEdge{To: 2, Weight: 5},
			graph.GraphEdge{To: 1, Weight: 1},
		},
		{
			graph.GraphEdge{To: 2, Weight: 7},
			graph.GraphEdge{To: 3, Weight: 6},
		},
		{graph.GraphEdge{To: 4, Weight: 1}},
		{graph.GraphEdge{To: 2, Weight: 1}},
		{},
	}

	expected := []int{0, 2, 4}
	actual := graph.DijkstraList(0, 4, graph_list)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
