package main

import (
	"fmt"

	graph "github.com/zoumas/dsa/go/graph/adjacency_matrix"
)

func main() {
	g := graph.NewGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 2)

	fmt.Println(g)
}
