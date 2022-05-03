package main

import "container/heap"

// Kruskal's algorithm for Minimum Spanning Tree
func kruskal(graph *Graph) (mst *Graph) {
	mst = NewGraph()
	h := &Heap{}
	heap.Init(h)
	// get a sorted list of edges
	for nodeName, edges := range graph.Edges {
		node := graph.GetNode(nodeName)
		for _, edge := range edges {
			heap.Push(h, NodePair{node, edge})
		}
		mst.AddNode(node)
	}

	for len(*h) > 0 {
		pair := heap.Pop(h).(NodePair)
		// if the edge isn't there between pair of nodes in the MST, creaet it
		if !mst.HasEdge(pair.node.name, pair.edge) {
			mst.AddEdge(pair.node, pair.edge.node, pair.edge.weight)
			// but if it creates a cyclical graph, remove it
			if mst.isCyclic(pair.node.name) {
				mst.RemoveEdge(pair.node.name, pair.edge.node.name)
			}
		}
	}

	return
}
