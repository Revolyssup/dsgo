package graph

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func New(n *Node) Graph {
	g := Graph{mx: &sync.Mutex{}, Log: *log.Default()}
	if n != nil {
		if n.Name == "" {
			name := randStringRunes(5)
			n.Name = name
			g.Log.Printf("Empty name detected. New name generated for this node: %s. Use this name for future references", name)
		}
		g.Nodes = append(g.Nodes, n)
	}
	return g
}

func (g *Graph) AddNode(n *Node, outTo map[string]Weight, inTo map[string]Weight) {
	g.mx.Lock()
	defer g.mx.Unlock()
	if g.Edge == nil {
		g.Edge = map[*Node]map[*Node]Weight{}
	}
	for _, node := range g.Nodes {
		if n.Name == node.Name {
			log.Println("could not add edge. An edge with same name exists")
			return
		}
	}

	for frm, wt := range inTo {
		for _, node := range g.Nodes {
			if node.Name == frm {
				if g.Edge[node] == nil {
					g.Edge[node] = make(map[*Node]Weight)
				}
				g.Edge[node][n] = wt
			}
		}
	}
	for to, wt := range outTo {
		for _, node := range g.Nodes {
			if node.Name == to {
				if g.Edge[n] == nil {
					g.Edge[n] = make(map[*Node]Weight)
				}
				g.Edge[n][node] = wt
			}
		}
	}
	g.Nodes = append(g.Nodes, n)
}
func (g *Graph) RemoveNode(name string) {
	g.mx.Lock()
	defer g.mx.Unlock()
	if name == "" {
		g.Log.Println("pass a non empty name")
		return
	}
	for fromedge := range g.Edge {
		if fromedge.Name == name {
			delete(g.Edge, fromedge)
			break
		}
	}

	for i, n := range g.Nodes {
		if n.Name == name {
			g.Nodes = append(g.Nodes[0:i], g.Nodes[i+1:]...)
			break
		}
	}
}
func (g *Graph) Print() {
	for _, n := range g.Nodes {
		g.Log.Printf("%s outgoing -->", n.Name)
		for n2, wt := range g.Edge[n] {
			g.Log.Printf("\t%s with weight %f\n", n2.Name, wt)
		}
	}
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
