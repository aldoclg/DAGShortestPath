package graph

type Edge struct {
	weight uint64
	target *Vertex
}

func NewEdge(vertex *Vertex, weight uint64) *Edge {
	return &Edge{weight: weight, target: vertex}
}

func (e *Edge) GetWeight() uint64 {
	return e.weight
}

func (e *Edge) GetTarget() *Vertex {
	return e.target
}
