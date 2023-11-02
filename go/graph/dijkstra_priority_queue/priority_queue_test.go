package graph_test

import (
	"math"
	"testing"

	graph "github.com/zoumas/dsa/go/graph/dijkstra_priority_queue"
)

func TestPriorityQueue(t *testing.T) {
	infinity := math.MaxInt

	q := new(graph.PriorityQueue[graph.VertexPriority])
	if !q.IsEmpty() {
		t.Fatal("Queue should have been empty")
	}

	q.Enqueue(graph.VertexPriority{1, infinity})
	q.Enqueue(graph.VertexPriority{2, 69})
	q.Enqueue(graph.VertexPriority{0, 0})

	v, ok := q.Dequeue()
	assertTrue(t, ok)
	assertPriority(t, v, 0)

	v, ok = q.Dequeue()
	assertTrue(t, ok)
	assertPriority(t, v, 69)

	v, ok = q.Dequeue()
	assertTrue(t, ok)
	assertPriority(t, v, infinity)
}

func assertTrue(t testing.TB, got bool) {
	t.Helper()

	if !got {
		t.Fatal("Not true")
	}
}

func assertPriority(t testing.TB, p graph.Prioritizer, want int) {
	t.Helper()

	if got := p.Priority(); got != want {
		t.Errorf("\nGot priority: %v\nWant priority: %v", got, want)
	}
}
