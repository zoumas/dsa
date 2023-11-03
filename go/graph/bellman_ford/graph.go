package graph

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
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

func (g Graph) BellmanFord(s int) *Graph {
	spt := NewGraph()

	// var infinity int
	// for _, v := range g.V() {
	// 	for _, n := range g.N(v) {
	// 		infinity += n.Weight
	// 	}
	// }
	const infinity = math.MaxInt32
	// Το πρόβλημα που είχα:
	// το math.MaxInt μου έκανε overflow το int και είχα
	// αρνητικό τεράστιο βάρος

	P := make(map[int]int)
	P[s] = s

	order := g.Order()
	B := make([]int, 0, order)
	Bnew := make([]int, 0, order)
	for i := 0; i < order; i++ {
		B = append(B, infinity)
		Bnew = append(Bnew, infinity)
	}
	B[s] = 0
	Bnew[s] = 0

	for count := 1; count <= g.Order()-1; count++ {
		for _, v := range g.V() {
			for _, n := range g.N(v) {
				if bn := B[n.Index] + n.Weight; bn < Bnew[v] {
					Bnew[v] = bn
					P[v] = n.Index
				}
			}
		}
		// memcpy
		for i := 0; i < order; i++ {
			B[i] = Bnew[i]
		}

		// Debug Print
		fmt.Println("Step:", count)
		fmt.Println("B:", B)
		fmt.Println("P:", P)
		fmt.Println("")
	}

	for v := 1; v < g.Order(); v++ {
		w, ok := g.EdgeWeight(v, P[v])
		if ok {
			spt.AddEdge(v, P[v], w)
		} else {
			log.Printf("Edge %v -> %v not found\n", v, P[v])
		}
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

func (g *Graph) AddEdge(u, v, w int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], Vertex{v, w})
	g.adjacencyList[v] = append(g.adjacencyList[v], Vertex{u, w})
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
			return nil, fmt.Errorf("Want 3 fields. Two vertices and an edge. Got %v", fields)
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
