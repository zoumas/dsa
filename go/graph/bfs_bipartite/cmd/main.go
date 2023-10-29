package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/bfs_bipartite"
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

	fmt.Println("Αρχικό Γράφημα")
	fmt.Println(g)

	ok, v1, v2 := g.IsBipartiteBFS(1)
	if ok {
		fmt.Println("Το γράφημα είναι διχοτομίσιμο")
		fmt.Println("v1:", v1)
		fmt.Println("v2:", v2)
	}
}
