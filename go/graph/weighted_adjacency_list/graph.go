package graph

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type Vertex struct {
	Index  int
	Weight int
}

type Graph struct {
	adjacencyList map[int][]Vertex
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]Vertex),
	}
}

func (g *Graph) AddEdge(u, v, weight int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], Vertex{Index: v, Weight: weight})
	g.adjacencyList[v] = append(g.adjacencyList[v], Vertex{Index: u, Weight: weight})
}

// Vertices returns a sorted list of the vertices of the Graph
func (g Graph) Vertices() []int {
	V := []int{}
	for k := range g.adjacencyList {
		V = append(V, k)
	}
	slices.Sort(V)
	return V
}

// Neighbors returns a sorted list of vertex's v Neighbors
func (g Graph) Neighbors(v int) []Vertex {
	sorted := make([]Vertex, len(g.adjacencyList[v]))
	copy(sorted, g.adjacencyList[v])

	slices.SortFunc(sorted, func(a, b Vertex) int {
		return cmp.Compare(a.Index, b.Index)
	})
	return sorted
}

// Degree returns the number of Neighbors that a vertex v has
func (g Graph) Degree(v int) int {
	return len(g.adjacencyList[v])
}

func (g Graph) String() string {
	var b strings.Builder
	V := g.Vertices()
	for _, v := range V {
		N := g.Neighbors(v)
		b.WriteString(fmt.Sprintf("N(%d): %+v\n", v, N))
	}
	return b.String()
}
