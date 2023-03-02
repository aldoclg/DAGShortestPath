package dag

import (
	"fmt"

	"github.com/aldoclg/DAGShortestPath/models/graph"
)

type ShortestPath struct {
	topologicalOrder *TopologicalOrder
}

func NewShortestPath(vertexes []*graph.Vertex) *ShortestPath {
	vertexes[0].MinDistance = 0
	shotestPath := &ShortestPath{topologicalOrder: NewTopologicalOrder(vertexes)}

	return shotestPath
}

func (s *ShortestPath) Compute() {
	topologicalOrderStack := s.topologicalOrder.GetStack()
	for topologicalOrderStack.IsNotEmpty() {
		u := topologicalOrderStack.Pop()
		for _, e := range u.GetAdjacency() {

			v := e.GetTarget()
			fmt.Printf("%v:%.1f - %.1f -> %v\n", u.GetName(), u.MinDistance, e.GetWeight(), v.GetName())
			shortestDistanceCandidate := u.MinDistance + e.GetWeight()
			if v.MinDistance > shortestDistanceCandidate {
				v.MinDistance = shortestDistanceCandidate
				v.Predecessor = u
			}
		}
	}
}
