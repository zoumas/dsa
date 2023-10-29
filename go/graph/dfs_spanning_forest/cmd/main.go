package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/dfs_spanning_tree"
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

	fmt.Println("Αρχικό Mη Συνδεδεμένο Γράφημα")
	fmt.Println(g)

	fmt.Println("Συνδετικό Δάσος της Διάσχισης κατά Βάθος")
	fmt.Print(g.SpanningForestDFS())
}
