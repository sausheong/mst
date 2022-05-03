package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortNodes(nodes []*Node) []*Node {
	sort.Slice(nodes, func(i, j int) bool {
		return strings.Compare(nodes[i].name, nodes[j].name) < 1
	})
	return nodes
}

func main() {
	graph := buildGraph()
	fmt.Println("GRAPH\n-----")
	nodes := sortNodes(graph.Nodes)
	for _, node := range nodes {
		fmt.Printf("%s -> %v\n", node.name, graph.Edges[node.name])
	}

	fmt.Println()
	bMST := boruvka(graph)
	fmt.Println("BORUVSKA MINIMUM SPANNING TREE\n-----------------------------")
	nodes = sortNodes(bMST.Nodes)
	for _, node := range nodes {
		fmt.Printf("%s -> %v\n", node.name, bMST.Edges[node.name])
	}

	fmt.Println()
	kMST := kruskal(graph)
	fmt.Println("KRUSKAL MINIMUM SPANNING TREE\n-----------------------------")
	nodes = sortNodes(kMST.Nodes)
	for _, node := range nodes {
		fmt.Printf("%s -> %v\n", node.name, kMST.Edges[node.name])
	}

	fmt.Println()
	pMST := prim(graph, "A")
	fmt.Println("PRIM MINIMUM SPANNING TREE\n-------------------------")
	nodes = sortNodes(pMST.Nodes)
	for _, node := range nodes {
		fmt.Printf("%s -> %v\n", node.name, pMST.Edges[node.name])
	}
}
