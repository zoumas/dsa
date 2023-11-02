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

func (g Graph) DijkstraShortestPathsTree(s int) *Graph {
	spt := NewGraph()

	infinity := math.MaxInt

	order := g.Order()
	L := make([]int, 0, order)
	Q := make(map[int]bool)
	P := make(map[int]int)

	for i := 0; i < order; i++ {
		L = append(L, infinity)
	}
	L[s] = 0
	P[s] = s

	count := 0
	for count < order {
		var v int

		minLength := infinity
		for _, u := range g.V() {
			if Q[u] {
				continue
			}

			if L[u] < minLength {
				minLength = L[u]
				v = u
			}
		}
		count++
		Q[v] = true

		w, ok := g.EdgeWeight(v, P[v])
		if ok {
			spt.AddEdge(P[v], v, w)
		}

		for _, n := range g.N(v) {
			if Q[n.Index] {
				continue
			}

			if l := L[v] + n.Weight; l < L[n.Index] {
				L[n.Index] = l
				P[n.Index] = v
			}
		}
	}

	for _, v := range g.V() {
		fmt.Printf("Συντομότερο μονοπάτι από %d εώς %d: ", s, v)
		sp := []int{}
		u := v
		for u != P[s] {
			sp = append(sp, u)
			u = P[u]
		}
		sp = append(sp, s)
		slices.Reverse(sp)
		fmt.Println(sp)
	}

	return spt
}

func (g Graph) EdgeWeight(u, v int) (w int, ok bool) {
	for _, n := range g.N(u) {
		if n.Index == v {
			return n.Weight, true
		}
	}
	for _, n := range g.N(v) {
		if n.Index == u {
			return n.Weight, true
		}
	}
	return 0, false
}

func (g Graph) Dijkstra(s, t int) []int {
	var infinity int
	for _, v := range g.V() {
		for _, n := range g.N(v) {
			infinity += n.Weight
		}
	}

	order := g.Order()
	L := make([]int, 0, order)
	Q := make(map[int]bool)
	P := make(map[int]int)

	for i := 0; i < order; i++ {
		L = append(L, infinity)
	}
	L[s] = 0
	P[s] = s

	count := 0
	for count < order {
		var v int

		minL := infinity
		for _, u := range g.V() {
			if Q[u] {
				continue
			}

			if L[u] < minL {
				minL = L[u]
				v = u
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

		b.WriteString(fmt.Sprintf("%v -> %+v\n", v, neighbors))
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
			return nil, fmt.Errorf("Need 3 fields. Two vertices and a weight. Got %v", fields)
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
