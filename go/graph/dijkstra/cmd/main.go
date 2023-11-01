package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/dijkstra"
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

	fmt.Println("Αρχικό Γράφημα G = (V, E, W)")
	fmt.Println(g)

	fmt.Println("Αλγόριθμος του Dijkstra")
	fmt.Println("Αφετηρία:   0")
	fmt.Println("Προορισμός: 7")
	fmt.Println("Συντομότερο Μονοπάτι:", g.Dijkstra(0, 7))
}
