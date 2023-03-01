package graph

type Edge struct {
	weight      uint64
	startVertex *Vertex
	target      *Vertex
}

func NewEdge(targetVertex *Vertex, weight uint64) *Edge {
	return &Edge{weight: weight, target: targetVertex}
}

func NewEdgeWithStartVertex(startVertex, targetVertex *Vertex, weight uint64) *Edge {
	return &Edge{weight: weight, startVertex: startVertex, target: targetVertex}
}

func (e *Edge) GetWeight() uint64 {
	return e.weight
}

func (e *Edge) GetTarget() *Vertex {
	return e.target
}

func (e *Edge) GetStartVertex() *Vertex {
	return e.startVertex
}
