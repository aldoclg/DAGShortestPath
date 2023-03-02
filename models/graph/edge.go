package graph

type Edge struct {
	weight      float64
	startVertex *Vertex
	target      *Vertex
}

func NewEdge(targetVertex *Vertex, weight float64) *Edge {
	return &Edge{weight: weight, target: targetVertex}
}

func NewEdgeWithStartVertex(startVertex, targetVertex *Vertex, weight float64) *Edge {
	return &Edge{weight: weight, startVertex: startVertex, target: targetVertex}
}

func (e *Edge) GetWeight() float64 {
	return e.weight
}

func (e *Edge) GetTarget() *Vertex {
	return e.target
}

func (e *Edge) GetStartVertex() *Vertex {
	return e.startVertex
}
