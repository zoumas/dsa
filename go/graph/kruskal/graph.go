package graph

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/ihebu/dsu"
)

type Vertex struct {
	Index  int
	Weight int
}

type Graph struct {
	adjacencyList map[int][]Vertex
}

type Edge struct {
	v, u, w int
}

func (g *Graph) Edges() []Edge {
	edges := []Edge{}

	for _, v := range g.V() {
		for _, n := range g.N(v) {
			if v < n.Index {
				edges = append(edges, Edge{v, n.Index, n.Weight})
			}
		}
	}

	return edges
}

func (g *Graph) Kruskal() *Graph {
	mst := NewGraph()

	d := dsu.New()
	for _, v := range g.V() {
		d.Add(v)
	}

	edges := g.Edges()
	slices.SortStableFunc(edges, func(a, b Edge) int {
		return cmp.Compare(a.w, b.w)
	})

	i, count := 0, 0
	for count < g.Order()-1 {
		if !reflect.DeepEqual(d.Find(edges[i].v), edges[i].u) {
			mst.AddEdge(edges[i].v, edges[i].u, edges[i].w)
			d.Union(edges[i].v, edges[i].u)
			count++
		}
		i++
	}

	return mst
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
