package main

import "sort"

// a subgraph
type Subgraph struct {
	nodes []*Node
	pairs []*NodePair
}

// Boruvka's algorithm for Minimum Spanning Tree
func boruvka(graph *Graph) (mst *Graph) {
	mst = NewGraph()
	// set up the MST
	for _, node := range graph.Nodes {
		mst.AddNode(node)
	}

	var subgraphs []*Subgraph
	// set up the subgraphs; intially each subgraph has only 1 node
	for _, node := range graph.Nodes {
		s := &Subgraph{nodes: []*Node{node}}
		for _, edge := range graph.Edges[node.name] {
			s.pairs = append(s.pairs, &NodePair{node, edge})
		}
		subgraphs = append(subgraphs, s)
	}

	// repeatedly combine the subgraphs until there is only 1
	for len(subgraphs) > 1 {
		pairs := subgraphs[0].pairs
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].edge.weight < pairs[j].edge.weight
		})
		mst.AddEdge(pairs[0].node, pairs[0].edge.node, pairs[0].edge.weight)
		subgraphs = combine(pairs[0].node, pairs[0].edge.node, subgraphs)
	}

	return
}

// combine 2 different subgraphs by connecting 2 nodes
func combine(n1, n2 *Node, subgraphs []*Subgraph) []*Subgraph {
	var s1, s2 *Subgraph
	var i2 int
	for i, subgraph := range subgraphs {
		// find the 2 subgraphs which has the 2 nodes
		for _, node := range subgraph.nodes {
			if n1.name == node.name {
				s1 = subgraph
			}
			if n2.name == node.name {
				s2 = subgraph
				i2 = i
			}
		}
	}
	// combine the nodes and node-pairs together
	s1.nodes = append(s1.nodes, s2.nodes...)
	s1.pairs = append(s1.pairs, s2.pairs...)

	// remove node-pairs that connect between 2 nodes that are in the same subgraph
	for _, subgraph := range subgraphs {
		var pairs []*NodePair
		for _, pair := range subgraph.pairs {
			if !(in(pair.node, subgraph) && in(pair.edge.node, subgraph)) {
				pairs = append(pairs, pair)
			}
		}
		subgraph.pairs = pairs
	}
	// remove subgraph s2
	subgraphs[i2] = subgraphs[len(subgraphs)-1]
	subgraphs = subgraphs[:len(subgraphs)-1]
	return subgraphs
}

// check if a node is in a subgraph
func in(node *Node, subgraph *Subgraph) bool {
	for _, n := range subgraph.nodes {
		if n.name == node.name {
			return true
		}
	}
	return false
}
