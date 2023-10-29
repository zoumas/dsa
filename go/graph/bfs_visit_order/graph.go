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
	keys := make([]int, 0, g.Order())
	for k := range g.adjacencyList {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func (g Graph) N(v int) []int {
	neighbors := slices.Clone(g.adjacencyList[v])
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

func (g Graph) VisitOrderBFS(start int) []int {
	visited := make(map[int]struct{})
	visited[start] = struct{}{}

	to_visit := []int{}
	to_visit = append(to_visit, start)

	visit_order := []int{start}

	for len(to_visit) != 0 {
		v := to_visit[0]
		to_visit = to_visit[1:]

		for _, n := range g.N(v) {
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				to_visit = append(to_visit, n)

				visit_order = append(visit_order, n)
			}
		}
	}

	return visit_order
}
