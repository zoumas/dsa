package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/bfs_spanning_tree"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	g, err := graph.NewGraphFromFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("G(V, E):")
	fmt.Print(g)

	fmt.Println()
	fmt.Println("Spanning Tree BFS")
	fmt.Print(g.SpanningTreeBFS(0))
}
