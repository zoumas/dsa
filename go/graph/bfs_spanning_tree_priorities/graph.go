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

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

func NewGraphFromFile(file *os.File) (*Graph, error) {
	g := NewGraph()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("Need 2 vertices to form an edge. Got %d", len(fields))
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

func (g *Graph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

func (g Graph) String() string {
	var b strings.Builder
	for _, v := range g.V() {
		b.WriteString(fmt.Sprintf("%d -> %v\n", v, g.Neighbors(v)))
	}
	return b.String()
}

// Επιστρέφει το Σύνολο κορυφών του γραφήματος ταξηνομημένο
func (g Graph) V() []int {
	vertices := []int{}
	for k := range g.adjacencyList {
		vertices = append(vertices, k)
	}
	slices.Sort(vertices)
	return vertices
}

// Σύνολο γειτόνων μία κορυφής v => N(v)
func (g Graph) Neighbors(v int) []int {
	s := slices.Clone(g.adjacencyList[v])
	slices.Sort(s)
	return s
}

// Τάξη του γραφήματος
func (g Graph) N() int {
	return len(g.adjacencyList)
}

func (g Graph) SpanningTreePrioritiesBFS(root int, priorities map[int]int) *Graph {
	spanningTree := NewGraph()

	visited := make(map[int]struct{}) // Mark visited/discovered vertices
	visited[root] = struct{}{}
	to_visit := []int{root} // Queue

	order := g.N()
	count := 1
	done := false
	for len(to_visit) != 0 && !done {
		v := to_visit[0]
		to_visit = to_visit[1:] // Dequeue

		neighbors := g.Neighbors(v)
		slices.SortStableFunc(neighbors, func(a, b int) int {
			return cmp.Compare(priorities[a], priorities[b])
		}) // sorts the neighbors by ascending priority order

		for _, n := range neighbors {
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				to_visit = append(to_visit, n) // Enqueue
				spanningTree.AddEdge(v, n)
				count++
				if count == order {
					done = true
					break
				}
			}
		}
	}

	return spanningTree
}
