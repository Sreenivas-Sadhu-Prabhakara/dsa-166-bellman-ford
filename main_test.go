package main

import (
	"reflect"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	// Negative edge shortens the path: 0->1->2 = 2 beats 0->2 = 4.
	dist, ok := BellmanFord(4, []Edge{{0, 1, 5}, {1, 2, -3}, {0, 2, 4}, {2, 3, 2}}, 0)
	if !ok || !reflect.DeepEqual(dist, []int{0, 5, 2, 4}) {
		t.Errorf("dist = %v, ok=%v", dist, ok)
	}

	// Reachable negative cycle: no answer.
	if _, ok := BellmanFord(3, []Edge{{0, 1, 1}, {1, 2, -2}, {2, 1, 1}}, 0); ok {
		t.Error("negative cycle must report not-ok")
	}

	// Unreachable negative cycle doesn't poison reachable answers.
	dist, ok = BellmanFord(4, []Edge{{2, 3, -2}, {3, 2, -2}, {0, 1, 7}}, 0)
	if !ok || dist[1] != 7 || dist[2] != -1 || dist[3] != -1 {
		t.Errorf("unreachable cycle: dist=%v ok=%v", dist, ok)
	}

	dist, ok = BellmanFord(1, nil, 0)
	if !ok || !reflect.DeepEqual(dist, []int{0}) {
		t.Errorf("single vertex = %v", dist)
	}
}
