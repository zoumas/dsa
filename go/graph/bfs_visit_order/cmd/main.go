package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/bfs_visit_order"
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

	fmt.Println("Ακολουθία κατά Πλάτος")
	fmt.Println(g.VisitOrderBFS(0))
}
