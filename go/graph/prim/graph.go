package graph

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Vertex struct {
	Index  int
	Weight int
}

type Graph struct {
	adjacencyList map[int][]Vertex
}

func (g *Graph) Prim() *Graph {
	mspt := NewGraph()

	var infinity int
	for _, v := range g.V() {
		for _, n := range g.N(v) {
			infinity += n.Weight
		}
	}

	// var start int
	// for _, v := range g.V() {
	// 	start = v
	// 	break
	// }
	// Start at a 'random' vertex
	start := 0

	type Edge struct {
		v, u int
	}
	E := make(map[int]Edge)

	Q := make(map[int]struct{})
	C := make(map[int]int)
	for _, v := range g.V() {
		Q[v] = struct{}{}
		C[v] = infinity
	}
	C[start] = 0

	for _, n := range g.N(start) {
		C[n.Index] = n.Weight
		E[n.Index] = Edge{start, n.Index}
	}
	delete(Q, start)

	for len(Q) != 0 {
		var v int

		minCost := infinity
		for vertex := range Q {
			if C[vertex] < minCost {
				minCost = C[vertex]
				v = vertex
			}
		}
		delete(Q, v)

		mspt.AddEdge(E[v].v, E[v].u, C[v])
		fmt.Fprintf(os.Stderr, "Added edge(%d, %d) with weight %d\n", E[v].v, E[v].u, C[v])

		for _, n := range g.N(v) {
			if _, ok := Q[n.Index]; !ok {
				continue
			}

			if n.Weight < C[n.Index] {
				C[n.Index] = n.Weight
				E[n.Index] = Edge{v, n.Index}
			}
		}
	}

	return mspt
}

func (g *Graph) AddEdge(u, v, w int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], Vertex{v, w})
	g.adjacencyList[v] = append(g.adjacencyList[v], Vertex{u, w})
}

func (g *Graph) Order() int {
	return len(g.adjacencyList)
}

func (g *Graph) V() []int {
	vertices := make([]int, 0, g.Order())
	for v := range g.adjacencyList {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g *Graph) N(v int) []Vertex {
	return g.adjacencyList[v]
}

func (g *Graph) String() string {
	var b strings.Builder

	vertices := g.V()
	slices.Sort(vertices)

	for _, v := range vertices {
		neighbors := g.N(v)
		slices.SortStableFunc(neighbors, func(a, b Vertex) int {
			return cmp.Compare(a.Index, b.Index)
		})

		b.WriteString(fmt.Sprintf("%v -> %v\n", v, neighbors))
	}

	return b.String()
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]Vertex),
	}
}

func NewGraphFromFile(file *os.File) (*Graph, error) {
	g := NewGraph()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := scanner.Text(), scanner.Err()
		if err != nil {
			return nil, err
		}

		fields := strings.Fields(line)
		if len(fields) != 3 {
			return nil, fmt.Errorf("Want 3 fields. Two vertices and a weight. Got %v", fields)
		}

		u, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		v, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		w, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, err
		}

		g.AddEdge(u, v, w)
	}

	return g, nil
}
