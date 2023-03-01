package graph

import (
	"fmt"
	"math"
)

type Vertex struct {
	name        string
	Visited     bool
	MinDistance uint64
	Predecessor *Vertex
	adjacency   []*Edge
}

type SortableByDistance []*Vertex

func NewVertex(name string) *Vertex {
	return &Vertex{name: name, MinDistance: math.MaxUint64, adjacency: make([]*Edge, 0)}
}

func (v *Vertex) AddNeighbor(edge *Edge) {
	v.adjacency = append(v.adjacency, edge)
}

func (v *Vertex) GetAdjacency() []*Edge {
	return v.adjacency
}

func (v *Vertex) GetName() string {
	return v.name
}

func (v *Vertex) ShowNeighbors() {
	for _, e := range v.adjacency {
		fmt.Println(e.target.name)
	}
}

func (v Vertex) String() string {
	o := fmt.Sprintf("%s", v.name)
	if v.Predecessor != nil {
		o = o + fmt.Sprintf(" - %v", v.Predecessor)
	}
	return o
}

func (v *Vertex) Compare(o *Vertex) bool {
	return v.MinDistance < o.MinDistance
}
