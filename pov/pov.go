package pov

import "fmt"

type Node string

type Graph map[Node][]Node

func New() *Graph {
	return &Graph{}
}

func (g *Graph) AddNode(nodeLabel string) {
	(*g)[Node(nodeLabel)] = []Node{}
}

func (g *Graph) AddArc(from, to string) {
	childs := (*g)[Node(from)]
	(*g)[Node(from)] = append(childs, Node(to))
}

func (g *Graph) ArcList() []string {
	edges := []string{}
	for from, nodes := range *g {
		for _, to := range nodes {
			edges = append(edges, fmt.Sprintf("%s -> %s", from, to))
		}
	}
	return edges
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	path := g.Path(Node(oldRoot), Node(newRoot))
	for i := 0; i < len(path)-1; i++ {
		oldTo, oldFrom := path[i], path[i+1]
		g.RemoveArc(oldFrom, oldTo)
		g.AddArc(string(oldTo), string(oldFrom))
	}
	return g
}

func (g *Graph) Path(from, to Node) []Node {
	if from == to {
		return []Node{to}
	}
	children := (*g)[from]
	for _, child := range children {
		if path := g.Path(child, to); path != nil {
			return append(path, from)
		}
	}
	return nil
}

func (g *Graph) RemoveArc(from, to Node) {
	children := (*g)[from]
	newChildren := []Node{}
	for _, child := range children {
		if child != to {
			newChildren = append(newChildren, child)
		}
	}
	(*g)[from] = newChildren
}
