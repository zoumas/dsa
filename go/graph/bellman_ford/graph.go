package graph

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
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

func (g *Graph) BellmanFord(start int) *Graph {
	spt := NewGraph()

	var infinity int
	for _, v := range g.V() {
		for _, n := range g.N(v) {
			infinity += n.Weight
		}
	}
	// Το άπειρο συμβολίζεται ως το άθροισμα όλων το βαρών του γραφήματος

	order := g.Order()

	B := make([]int, 0, order)
	Bnew := make([]int, 0, order)
	for i := 0; i < order; i++ {
		B = append(B, infinity)
		Bnew = append(Bnew, infinity)
	}
	B[start] = 0
	Bnew[start] = 0

	// Η αρχικοποίηση του B γίνεται όπως και ο πίνακας ετικετών L του αλγόριθμου του Dijkstra
	// Για όλες τις κορυφές έχουμε άπειρο εκτός της αφετηρίας που έχει μηδέν.
	// Bnew = B

	P := make(map[int]int)
	P[start] = start

	for i := 1; i < order; i++ {
		for _, v := range g.V() {
			for _, n := range g.N(v) {
				if b := B[n.Index] + n.Weight; b < Bnew[v] {
					Bnew[v] = b
					P[v] = n.Index
				}
			}
		}
		copy(B, Bnew)

		// Debug Print
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Step:", i)
		fmt.Fprintln(os.Stderr, "B:", B)
		fmt.Fprintln(os.Stderr, "P:", P)
	}

	for v := 1; v < order; v++ {
		w, ok := g.Weight(v, P[v])
		if ok {
			spt.AddEdge(v, P[v], w)
		} else {
			log.Printf("Unable to find weight between %v and %v\n", v, P[v])
		}
	}

	return spt
}

func (g *Graph) Weight(u, v int) (w int, ok bool) {
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
