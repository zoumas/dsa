package main

import (
	"fmt"

	graph "github.com/zoumas/dsa/go/graph/edges"
)

func main() {
	g := graph.NewGraph(4)

	g.AddEdge(1, 0, 5)
	g.AddEdge(0, 2, 3)
	g.AddEdge(0, 3, 1)
	g.AddEdge(2, 1, 2)

	fmt.Println(g)
	fmt.Println("|E| = ", g.Size())
}
