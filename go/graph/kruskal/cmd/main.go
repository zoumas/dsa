package main

import (
	"fmt"
	"log"
	"os"

	graph "github.com/zoumas/dsa/go/graph/kruskal"
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

	fmt.Println("Kruskal's Minimum Spanning Tree")
	mst := g.Kruskal()
	fmt.Println(mst)

	weight := 0
	for _, v := range mst.V() {
		for _, n := range mst.N(v) {
			weight += n.Weight
		}
	}
	fmt.Println("Weight:", weight)
}
