package main

import "container/heap"

// Prim's algorithm for Minimum Spanning Tree
func prim(graph *Graph, nodeName string) (mst *Graph) {
	mst = NewGraph()
	h := &Heap{}
	heap.Init(h)

	startNode := graph.GetNode(nodeName)
	mst.AddNode(startNode)
	for _, edge := range graph.Edges[startNode.name] {
		heap.Push(h, NodePair{startNode, edge})
	}

	for len(*h) > 0 {
		// get the edge with the smallest weight
		pair := heap.Pop(h).(NodePair)
		// skip if the node pair is already in the MST
		if !mst.HasNode(pair.edge.node.name) {
			// add the node and the edge to the MST
			mst.AddNode(pair.edge.node)
			mst.AddEdge(pair.node, pair.edge.node, pair.edge.weight)
			// add all the edges going from the connected node into the heap
			for _, edge := range graph.Edges[pair.edge.node.name] {
				heap.Push(h, NodePair{pair.edge.node, edge})
			}
		}
	}
	return
}
