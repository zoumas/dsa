package graph

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

func NewGraphFromFile(file *os.File) (*Graph, error) {
	g := NewGraph()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if l := len(fields); l != 2 {
			return nil, fmt.Errorf("Need 2 vertices to form an edge. Got %d", l)
		}
		u, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		v, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}

		g.AddEdge(u, v)
	}

	return g, nil
}

func (g Graph) Order() int {
	return len(g.adjacencyList)
}

func (g Graph) V() []int {
	vertices := []int{}
	for v := range g.adjacencyList {
		vertices = append(vertices, v)
	}
	slices.Sort(vertices)
	return vertices
}

func (g Graph) N(v int) []int {
	neighbors := g.adjacencyList[v]
	slices.Sort(neighbors)
	return neighbors
}

func (g Graph) String() string {
	var b strings.Builder
	for _, v := range g.V() {
		b.WriteString(fmt.Sprintf("%d -> %v\n", v, g.N(v)))
	}
	return b.String()
}

type Edge struct {
	v      int
	parent int
}

// Θα μπορούσα να τρέχω SpanningTreeDFS για κάθε συνιστώσα.
func (g Graph) SpanningForestDFS() *Graph {
	forest := NewGraph()

	discovered := make(map[int]struct{})

	order := g.Order()
	count := 0
	for {
		var start int
		if count >= order {
			break
		}

		for _, v := range g.V() {
			if _, ok := discovered[v]; !ok {
				start = v
				break
			}
		}

		stack := []Edge{}
		discovered[start] = struct{}{}
		count++

		neighbors := g.N(start)
		slices.Reverse(neighbors)
		for _, n := range neighbors {
			stack = append(stack, Edge{v: n, parent: start})
		}

		for len(stack) != 0 {
			top := len(stack) - 1
			e := stack[top]
			stack = stack[:top]

			if _, ok := discovered[e.v]; !ok {
				discovered[e.v] = struct{}{}
				forest.AddEdge(e.parent, e.v)
				count++

				neighbors := g.N(e.v)
				slices.Reverse(neighbors)
				for _, n := range neighbors {
					stack = append(stack, Edge{v: n, parent: e.v})
				}
			}
		}
	}

	return forest
}
