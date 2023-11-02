package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/dijkstra_shortest_path_tree"
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

	fmt.Println("Αρχικό Γράφημα")
	fmt.Println(g)
	fmt.Println("Αλγόριθμος του Dijkstra")
	spt := g.DijkstraShortestPathsTree(0)
	fmt.Println("Δένδρο Συντομότερων Μονοπατιών")
	fmt.Print(spt)
}
