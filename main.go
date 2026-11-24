package main

import "fmt"

// Edge is a directed weighted edge.
type Edge struct {
	U, V, W int
}

// BellmanFord returns shortest distances from src, allowing negative edge
// weights. ok is false when a negative cycle is reachable from src (no
// well-defined distances). Unreachable vertices carry dist -1 and ok stays
// true. V-1 relaxation rounds suffice; a V-th improvement betrays the cycle.
func BellmanFord(n int, edges []Edge, src int) (dist []int, ok bool) {
	const inf = int(^uint(0) >> 1)
	dist = make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0
	for round := 0; round < n-1; round++ {
		changed := false
		for _, e := range edges {
			if dist[e.U] != inf && dist[e.U]+e.W < dist[e.V] {
				dist[e.V] = dist[e.U] + e.W
				changed = true
			}
		}
		if !changed {
			break // already settled; later rounds can't help
		}
	}
	for _, e := range edges {
		if dist[e.U] != inf && dist[e.U]+e.W < dist[e.V] {
			return nil, false // still improvable: negative cycle
		}
	}
	for i, d := range dist {
		if d == inf {
			dist[i] = -1
		}
	}
	return dist, true
}

func main() {
	edges := []Edge{{0, 1, 5}, {1, 2, -3}, {0, 2, 4}, {2, 3, 2}}
	fmt.Println(BellmanFord(4, edges, 0))
}
