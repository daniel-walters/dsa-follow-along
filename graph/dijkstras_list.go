package graph

import "math"

func DijkstraList(source, sink int, arr WeightedAdjacencyList) []int {
	seen := []bool{}
	dists := []int{}
	prev := []int{}

	for range arr {
		seen = append(seen, false)
		dists = append(dists, math.MaxInt)
		prev = append(prev, -1)
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := arr[curr]

		for _, edge := range adjs {
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	out := []int{}
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)

	return reverse(out)
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i, isSeen := range seen {
		if !isSeen && dists[i] < math.MaxInt {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDist := math.MaxInt

	for i, isSeen := range seen {
		if isSeen {
			continue
		}

		if lowestDist > dists[i] {
			lowestDist = dists[i]
			idx = i
		}
	}

	return idx
}
