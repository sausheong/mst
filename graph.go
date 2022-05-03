package main

import (
	"fmt"
	"strconv"
)

func buildGraph() *Graph {
	graph := NewGraph()
	nodes := make(map[string]*Node)
	for _, name := range []string{"A", "B", "C", "D", "E", "F", "G"} {
		n := &Node{name, 0}
		graph.AddNode(n)
		nodes[name] = n
	}

	graph.AddEdge(nodes["A"], nodes["B"], 7)
	graph.AddEdge(nodes["A"], nodes["D"], 4)
	graph.AddEdge(nodes["B"], nodes["C"], 11)
	graph.AddEdge(nodes["B"], nodes["D"], 9)
	graph.AddEdge(nodes["B"], nodes["E"], 10)
	graph.AddEdge(nodes["C"], nodes["E"], 5)
	graph.AddEdge(nodes["D"], nodes["E"], 15)
	graph.AddEdge(nodes["D"], nodes["F"], 6)
	graph.AddEdge(nodes["E"], nodes["F"], 12)
	graph.AddEdge(nodes["E"], nodes["G"], 8)
	graph.AddEdge(nodes["F"], nodes["G"], 13)

	return graph
}

// -- Weighted Undirected Graph

type Node struct {
	name  string
	value int
}

func (n *Node) String() string {
	return n.name
}

type Edge struct {
	node   *Node
	weight int
}

func (e *Edge) String() string {
	return e.node.String() + "(" + strconv.Itoa(e.weight) + ")"
}

type Graph struct {
	Nodes []*Node
	Edges map[string][]*Edge // key is node name
}

func (g *Graph) String() (s string) {
	for _, n := range g.Nodes {
		s = s + n.String() + " ->"
		for _, c := range g.Edges[n.name] {
			s = s + " " + c.node.String() + " (" + strconv.Itoa(c.weight) + ")"
		}
		s = s + "\n"
	}
	return
}

type NodePair struct {
	node *Node
	edge *Edge
}

func (np *NodePair) String() string {
	return fmt.Sprintf("%s -> %s (%d)", np.node.name, np.edge.node.name, np.edge.weight)
}

func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[string][]*Edge),
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{n2, weight})
	g.Edges[n2.name] = append(g.Edges[n2.name], &Edge{n1, weight})
}

func (g *Graph) GetNode(name string) (node *Node) {
	for _, n := range g.Nodes {
		if n.name == name {
			node = n
		}
	}
	return
}

func (g *Graph) GetEdges(n1, n2 string) (edges []*Edge) {
	for _, edge := range g.Edges[n1] {
		if edge.node.name == n2 {
			edges = append(edges, edge)
		}
	}
	for _, edge := range g.Edges[n2] {
		if edge.node.name == n1 {
			edges = append(edges, edge)
		}
	}
	return
}

func (g *Graph) HasNode(name string) (yes bool) {
	for _, n := range g.Nodes {
		if n.name == name {
			yes = true
		}
	}
	return
}

func (g *Graph) HasEdge(name string, edge *Edge) bool {
	for _, e := range g.Edges[name] {
		if e.node.name == edge.node.name && e.weight == edge.weight {
			return true
		}
	}
	return false
}

func (g *Graph) RemoveNode(name string) {
	r := -1
	for i, n := range g.Nodes {
		if n.name == name {
			r = i
		}
	}
	if r > -1 {
		g.Nodes[r] = g.Nodes[len(g.Nodes)-1] // remove the node
		g.Nodes = g.Nodes[:len(g.Nodes)-1]
	}
	delete(g.Edges, name) // remove the edge from one side
	// remove the edge from the other side
	for n := range g.Edges {
		rmEdge(g, n, name)

	}
}

func (g *Graph) RemoveEdge(n1, n2 string) {
	rmEdge(g, n1, n2)
	rmEdge(g, n2, n1)
}

func rmEdge(g *Graph, m, n string) {
	edges := g.Edges[m]
	r := -1
	for i, edge := range edges {
		if edge.node.name == n {
			r = i
		}
	}
	if r > -1 {
		edges[r] = edges[len(edges)-1]
		g.Edges[m] = edges[:len(edges)-1]
	}
}

func (g *Graph) isCyclic(name string) (yes bool) {
	stack := &Stack{}
	visited := make(map[string]bool)
	startNode := g.GetNode(name)
	// use a 2 node array with first node as the from node and the second as the to node
	stack.Push([2]*Node{startNode, startNode})

	for len(stack.nodes) > 0 {
		n := stack.Pop()
		// if it's not visited before
		if !visited[n[1].name] {
			visited[n[1].name] = true // set to visited
			// get edges
			edges := g.Edges[n[1].name]
			for _, edge := range edges {
				// if this is not the from node
				if edge.node.name != n[0].name {
					// if it's visited the graph is cyclical
					if visited[edge.node.name] {
						return true
					}
					stack.Push([2]*Node{n[1], edge.node})
				}
			}
		}
	}
	return
}

// -- Min Heap for NodePair

type Heap []NodePair

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].edge.weight < h[j].edge.weight }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(NodePair)) }

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// -- Stack

type Stack struct {
	nodes [][2]*Node
}

func (s *Stack) Push(n [2]*Node) {
	s.nodes = append(s.nodes, n)
}

func (s *Stack) Pop() (n [2]*Node) {
	n = s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return
}
