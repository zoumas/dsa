package main

import (
	"fmt"

	graph "github.com/zoumas/dsa/go/graph/directed_adjacency_matrix"
)

func main() {
	g := graph.NewGraph(4)

	g.AddEdge(0, 3)
	g.AddEdge(0, 2)
	g.AddEdge(1, 0)
	g.AddEdge(2, 1)

	fmt.Println(g)
}
