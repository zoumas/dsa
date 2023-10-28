package main

import (
	"fmt"

	graph "github.com/zoumas/dsa/go/graph/directed_incidence_matrix"
)

func main() {
	g := graph.NewGraph(4, 4)

	g.AddEdge(0, 0, 3)
	g.AddEdge(1, 0, 1)
	g.AddEdge(2, 0, 2)
	g.AddEdge(3, 1, 2)

	fmt.Println(g)
}
