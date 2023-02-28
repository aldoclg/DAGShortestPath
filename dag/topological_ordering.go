package dag

import (
	"github.com/aldoclg/DAGShortestPath/models/graph"
	"github.com/aldoclg/DAGShortestPath/models/stack"
)

type TopologicalOrder struct {
	stack stack.Stack[graph.Vertex]
}

func NewTopologicalOrder(vertexes []*graph.Vertex) *TopologicalOrder {
	stack := stack.NewStack[graph.Vertex]()
	o := &TopologicalOrder{stack: *stack}

	for _, v := range vertexes {
		if !v.Visited {
			o.dfs(v)
		}
	}
	return o
}

func (to *TopologicalOrder) dfs(vertex *graph.Vertex) {
	vertex.Visited = true
	for _, e := range vertex.GetAdjacency() {
		to.dfs(e.GetTarget())
	}
	to.stack.Push(vertex)
}

func (to *TopologicalOrder) GetStack() *stack.Stack[graph.Vertex] {
	return &to.stack
}
