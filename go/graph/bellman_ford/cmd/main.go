package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/bellman_ford"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	g, err := graph.NewGraphFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Initial Graph")
	fmt.Println(g)

	fmt.Println("Bellman-Ford Shortest Paths Tree")
	fmt.Println(g.BellmanFord(0))
}
