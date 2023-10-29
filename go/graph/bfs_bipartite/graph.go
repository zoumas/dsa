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
	vertices := make([]int, 0, g.Order())
	for k := range g.adjacencyList {
		vertices = append(vertices, k)
	}
	slices.Sort(vertices)
	return vertices
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

const (
	partOfV1 = 1
	partOfV2 = 2
)

// Έλεγχος αν ένα γράφημα είναι διχοτομίσιμο με ΔκΠ
func (g Graph) IsBipartiteBFS(start int) (isBipartite bool, v1, v2 []int) {
	visited := make(map[int]int)
	to_visit := []int{start}

	visited[start] = partOfV1
	v1 = append(v1, start)

	for len(to_visit) != 0 {
		// dequeue
		v := to_visit[0]
		to_visit = to_visit[1:]

		for _, n := range g.N(v) {
			if _, ok := visited[n]; !ok {
				if visited[v] == partOfV1 {
					visited[n] = partOfV2
					v2 = append(v2, n)
				} else {
					visited[n] = partOfV1
					v1 = append(v1, n)
				}

				// enqueue
				to_visit = append(to_visit, n)
			} else if visited[n] == visited[v] {
				return false, nil, nil
			}
		}
	}

	return true, v1, v2
}
