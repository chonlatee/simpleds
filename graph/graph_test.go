package graph

import "testing"

func TestNewDirectedGraph(t *testing.T) {
	g := NewDirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 2)

	g.AddEdge(2, 3)

	g.AddEdge(2, 4)

	g.AddEdge(4, 1)
}
