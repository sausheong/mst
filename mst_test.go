package main

import "testing"

func BenchmarkBoruvka(b *testing.B) {
	graph := buildGraph()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boruvka(graph)
	}

}

func BenchmarkKruskal(b *testing.B) {
	graph := buildGraph()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kruskal(graph)
	}
}

func BenchmarkPrim(b *testing.B) {
	graph := buildGraph()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prim(graph, "A")
	}
}
