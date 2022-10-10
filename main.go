package main

import (
	"fmt"

	"github.com/Revolyssup/dsgo/graph"
)

func main() {
	g := graph.New(&graph.Node{
		Data: graph.Data{
			Value: 23,
		},
		Name: "A",
	})
	fill(&g)
	fmt.Println(g.DijkstraShortestPathWeight("A", "F"))
}

func fill(g *graph.Graph) {
	// var bout = make(map[string]graph.Weight)
	var bin = make(map[string]graph.Weight)
	bin["A"] = 1
	g.AddNode(&graph.Node{
		Name: "B",
		Data: graph.Data{
			Value: 26,
		},
	}, nil, bin)
	delete(bin, "A")
	bin["B"] = 1
	g.AddNode(&graph.Node{
		Name: "C",
		Data: graph.Data{
			Value: 26,
		},
	}, nil, bin)
	delete(bin, "B")

	bin["A"] = 1
	g.AddNode(&graph.Node{
		Name: "D",
		Data: graph.Data{
			Value: 26,
		},
	}, nil, bin)
	delete(bin, "A")

	bin["D"] = 1
	// bin["D"] = 1
	// g.AddNode(&graph.Node{
	// 	Name: "E",
	// 	Data: graph.Data{
	// 		Value: 26,
	// 	},
	// }, nil, bin)
	// delete(bin, "D")
	// bin["E"] = 1
	// bin["C"] = 1
	g.AddNode(&graph.Node{
		Name: "F",
		Data: graph.Data{
			Value: 26,
		},
	}, nil, bin)
	g.AddNode(&graph.Node{
		Name: "X",
		Data: graph.Data{
			Value: "unconnected",
		},
	}, nil, nil)
	g.Print()
}
