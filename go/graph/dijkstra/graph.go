package graph

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
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

func (g Graph) Dijkstra(s, t int) []int {
	// var infinity int
	// for _, v := range g.V() {
	// 	for _, n := range g.N(v) {
	// 		infinity += n.Weight
	// 	}
	// }
	const infinity = math.MaxInt
	order := g.Order()

	Q := make(map[int]bool)
	L := make([]int, 0, order)
	P := make(map[int]int)

	for i := 0; i < order; i++ {
		L = append(L, infinity)
	}
	L[s] = 0
	P[s] = s

	count := 0
	for count < order {
		var v int

		min := infinity
		for i := 0; i < order; i++ {
			if Q[i] {
				continue
			}

			if L[i] < min {
				min = L[i]
				v = i
			}
		}
		count++
		Q[v] = true

		if v == t {
			break
		}

		for _, n := range g.N(v) {
			if Q[n.Index] {
				continue
			}

			if L[v]+n.Weight < L[n.Index] {
				L[n.Index] = L[v] + n.Weight
				P[n.Index] = v
			}
		}
	}

	sp := []int{}
	v := t
	for v != P[s] {
		sp = append(sp, v)
		v = P[v]
	}
	sp = append(sp, s)
	slices.Reverse(sp)
	return sp
}

func (g *Graph) AddEdge(u, v, w int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], Vertex{v, w})
	g.adjacencyList[v] = append(g.adjacencyList[v], Vertex{u, w})
}

func (g Graph) Order() int {
	return len(g.adjacencyList)
}

func (g Graph) V() []int {
	vertices := make([]int, 0, g.Order())
	for v := range g.adjacencyList {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g Graph) N(v int) []Vertex {
	return g.adjacencyList[v]
}

func (g Graph) String() string {
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
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 3 {
			return nil, fmt.Errorf("Didn't get 3 fields. Got %d", len(fields))
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
