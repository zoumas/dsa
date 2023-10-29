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

func (g *Graph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

func (g Graph) Neighbors(v int) []int {
	N := slices.Clone(g.adjacencyList[v])
	slices.Sort(N)
	return N
}

func (g Graph) V() []int {
	keys := []int{}
	for k := range g.adjacencyList {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func (g Graph) Order() int {
	return len(g.adjacencyList)
}

func (g Graph) String() string {
	var b strings.Builder

	for _, v := range g.V() {
		b.WriteString(fmt.Sprintf("%d -> %v\n", v, g.Neighbors(v)))
	}
	return b.String()
}

func (g Graph) ShortestPathBFS(start, target int) []int {
	visited := make(map[int]struct{})
	to_visit := []int{}          // queue
	parent := make(map[int]*int) // shortest path lookup table

	to_visit = append(to_visit, start)
	visited[start] = struct{}{}
	parent[start] = &start

	done := false
	for len(to_visit) != 0 && !done {
		v := to_visit[0]
		to_visit = to_visit[1:] // dequeue

		for _, n := range g.Neighbors(v) {
			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				parent[n] = &v

				to_visit = append(to_visit, n)
				if target == v {
					done = true
					break
				}
			}
		}
	}

	path := []int{}
	v := target
	for parent[v] != &start {
		path = append([]int{v}, path...)
		v = *parent[v]
	}
	path = append([]int{start}, path...)

	return path
}
