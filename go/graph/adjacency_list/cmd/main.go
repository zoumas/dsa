package main

import (
	"fmt"

	graph "github.com/zoumas/dsa/go/graph/adjacency_list"
)

func main() {
	g := graph.NewGraph()

	g.AddEdge(1, 4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	fmt.Print(g)
	fmt.Println()

	for _, v := range g.Vertices() {
		fmt.Printf("Βαθμός Κορυφής d(%d): %d\n", v, g.Degree(v))
	}
}
