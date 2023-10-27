package graph

import (
	"fmt"
	"strings"
)

// Fixed Vertex Size Undirected Weighted Graph represented with an Adjacency Matrix
type Graph struct {
	n               int
	adjacencyMatrix [][]int
}

func NewGraph(vertices int) *Graph {
	matrix := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		matrix[i] = make([]int, vertices)
	}
	return &Graph{
		n:               vertices,
		adjacencyMatrix: matrix,
	}
}

func (g Graph) indexIsOutOfBounds(i int) bool {
	return i < 0 || i >= g.n
}

func (g *Graph) AddEdge(u, v, weight int) error {
	if g.indexIsOutOfBounds(u) {
		return fmt.Errorf("vertex index %d is out of bounds", u)
	}
	if g.indexIsOutOfBounds(v) {
		return fmt.Errorf("vertex index %d is out of bounds", v)
	}

	g.adjacencyMatrix[u][v] = weight
	g.adjacencyMatrix[v][u] = weight

	return nil
}

// Size returns the number of edges in the graph
func (g Graph) Size() int {
	count := 0
	for i := 0; i < g.n; i++ {
		for j := i + 1; j < g.n; j++ {
			if g.adjacencyMatrix[i][j] != 0 {
				count++
			}
		}
	}
	return count
}

func (g Graph) String() string {
	var b strings.Builder
	last := g.n - 1
	for i := 0; i < last; i++ {
		b.WriteString(fmt.Sprintln(g.adjacencyMatrix[i]))
	}
	b.WriteString(fmt.Sprint(g.adjacencyMatrix[last]))
	return b.String()
}
