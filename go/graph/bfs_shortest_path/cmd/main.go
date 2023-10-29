package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/bfs_shortest_path"
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
	fmt.Println(g)

	fmt.Println("Shortest Path from 0 to 7")
	fmt.Println(g.ShortestPathBFS(0, 7))
}
