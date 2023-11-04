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

	fmt.Println("G = (V, E, W)")
	fmt.Println(g)

	start := 0
	fmt.Println("Bellman Ford Shortest Paths Tree")
	fmt.Println("Start:", start)
	spt := g.BellmanFord(start)
	fmt.Println("")
	fmt.Print(spt)
}
