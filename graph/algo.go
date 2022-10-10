package graph

import (
	"fmt"
	"math"
)

func (g *Graph) DijkstraShortestPathWeight(from string, to string) (w Weight, err error) {
	var fromNode *Node
	var toNode *Node
	for _, node := range g.Nodes {
		if node.Name == from {
			fromNode = node
		}
		if node.Name == to {
			toNode = node
		}
	}
	if fromNode == nil {
		err = fmt.Errorf("no such fromnode")
		return
	}
	if toNode == nil {
		err = fmt.Errorf("no such tonode")
		return
	}
	processed := make(map[*Node]bool)
	cost := make(map[*Node]Weight)

	for _, n := range g.Nodes {
		if g.Edge[fromNode][n] == 0 && n != fromNode {
			cost[n] = math.MaxFloat64
		} else {
			cost[n] = g.Edge[fromNode][n]
		}
	}
	g._processNode(processed, fromNode, cost)
	return cost[toNode], err
}

func (g *Graph) _processNode(processed map[*Node]bool, node *Node, cost map[*Node]Weight) {
	if processed[node] {
		return
	}
	for neighbor, wt := range g.Edge[node] {
		if wt+cost[node] < cost[neighbor] {
			cost[neighbor] = wt + cost[node]
		}
	}
	processed[node] = true
	for neighbor := range g.Edge[node] {
		g._processNode(processed, neighbor, cost)
	}
}
