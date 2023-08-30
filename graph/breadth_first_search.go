package graph

import (
	"dsa/structures"
	"errors"
)

type WeightedAdjacencyMatrix = [][]int

func BFS(graph WeightedAdjacencyMatrix, source, needle int) ([]int, error) {
	seen := []bool{}
	prev := []int{}

	for range graph {
		seen = append(seen, false)
		prev = append(prev, -1)
	}

	q := structures.NewQueue[int]()

	seen[source] = true
	q.Enqueue(source)

	for q.Length() > 0 {
		curr, err := q.Dequeue()
		if err != nil {
			panic(err)
		}

		seen[curr] = true

		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := range adjs {
			if adjs[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q.Enqueue(i)
		}
	}

	// build it backwards
	curr := needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) > 0 {
		out = append(out, source)
		return reverse(out), nil
	}

	return nil, errors.New("Could not find a path from source to needle")
}

func reverse(arr []int) []int {
	out := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		out = append(out, arr[i])
	}

	return out
}
