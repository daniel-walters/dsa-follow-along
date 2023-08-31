package graph

import "errors"

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]GraphEdge

func walk(
	graph WeightedAdjacencyList,
	curr,
	needle int,
	seen *[]bool,
	path *[]int) bool {

	if (*seen)[curr] {
		return false
	}

	(*seen)[curr] = true

	//pre
	*path = append(*path, curr)
	if curr == needle {
		return true
	}

	//recurse
	list := graph[curr]
	for i := range list {
		edge := list[i]

		if walk(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	//post
	*path = (*path)[:len(*path)-1]

	return false
}

func DFS(
	graph WeightedAdjacencyList,
	source,
	needle int) ([]int, error) {
        seen := []bool{}
        path := []int{}

        for range graph {
            seen = append(seen, false)
        }

        walk(graph, source, needle, &seen, &path)

        if len(path) == 0 {
            return nil, errors.New("No path found from source to needle")
        }

	return path, nil
}
