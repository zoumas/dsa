package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/bfs_spanning_tree_priorities"
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
	fmt.Println("Spanning Tree BFS with Priorities")
	priorities := map[int]int{
		0:  0,
		1:  8,
		2:  4,
		3:  9,
		4:  1,
		5:  3,
		6:  10,
		7:  2,
		8:  7,
		9:  5,
		10: 6,
	}

	fmt.Print(g.SpanningTreePrioritiesBFS(4, priorities))
}
