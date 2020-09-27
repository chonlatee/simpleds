package graph

type Vertex struct {
	Key int

	Edge []*Vertex
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:  key,
		Edge: []*Vertex{},
	}
}

type Graph struct {
	Vertices map[int]*Vertex

	directed bool
}

func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]
	if !contain(v1.Edge, v2) {
		v1.Edge = append(v1.Edge, v2)
	}
}

func (g *Graph) TraverseDFS() {

}

func contain(vertices []*Vertex, vt *Vertex) bool {
	for _, v := range vertices {
		if v.Key == vt.Key {
			return true
		}
	}

	return false
}
