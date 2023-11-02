package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/dijkstra_priority_queue"
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
	fmt.Println("Συντομότερο μονοπάτι από 0 σε 7:", g.Dijkstra(0, 7))

	fmt.Println("Δένδρο Συντομότερων Μονοπατιών με αφετηρία την Κορυφή 0")
	fmt.Println(g.DijkstraShortestPathTree(0))
}
