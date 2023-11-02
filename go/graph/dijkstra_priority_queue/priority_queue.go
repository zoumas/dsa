package graph

import (
	"cmp"
	"slices"
)

type Prioritizer interface {
	Priority() int
}

type VertexPriority struct {
	Index    int
	Distance int
}

func (vp VertexPriority) Priority() int {
	return vp.Distance
}

type PriorityQueue[T Prioritizer] struct {
	buffer []T
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return len(q.buffer) == 0
}

func (q *PriorityQueue[T]) Enqueue(v T) {
	q.buffer = append(q.buffer, v)
	// sort in ascending priority order
	// refactor with an insertion sort later
	slices.SortStableFunc(q.buffer, func(a, b T) int {
		return cmp.Compare(a.Priority(), b.Priority())
	})
}

func (q *PriorityQueue[T]) Dequeue() (v T, ok bool) {
	if q.IsEmpty() {
		return v, false
	}

	v = q.buffer[0]
	q.buffer = q.buffer[1:]

	return v, true
}
