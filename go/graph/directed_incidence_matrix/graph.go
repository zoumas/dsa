package graph

import (
	"fmt"
	"strings"
)

type Graph struct {
	n, m            int
	incidenceMatrix [][]int
}

func NewGraph(vertices, edges int) *Graph {
	matrix := make([][]int, vertices)
	for i := 0; i < edges; i++ {
		matrix[i] = make([]int, edges)
	}
	return &Graph{
		n:               vertices,
		m:               edges,
		incidenceMatrix: matrix,
	}
}

func (g *Graph) AddEdge(edge, u, v int) error {
	if edge < 0 || edge >= g.m {
		return fmt.Errorf("edge index '%d' is out of bounds", edge)
	}

	if u < 0 || u >= g.n {
		return fmt.Errorf("vertex index '%d' is out of bounds", u)
	}
	if v < 0 || v >= g.n {
		return fmt.Errorf("vertex index '%d' is out of bounds", v)
	}

	g.incidenceMatrix[u][edge] = 1
	g.incidenceMatrix[v][edge] = -1

	return nil
}

func (g Graph) String() string {
	var b strings.Builder
	last := g.n - 1
	for i := 0; i < last; i++ {
		for j := 0; j < g.m; j++ {
			b.WriteString(fmt.Sprintf("%2d ", g.incidenceMatrix[i][j]))
		}
		b.WriteByte('\n')
	}
	for j := 0; j < g.m; j++ {
		b.WriteString(fmt.Sprintf("%2d ", g.incidenceMatrix[last][j]))
	}
	return b.String()
}
