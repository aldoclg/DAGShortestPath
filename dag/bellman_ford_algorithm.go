package dag

import (
	"fmt"
	"math"

	"github.com/aldoclg/DAGShortestPath/models/graph"
)

type BellmanFordAlgorithm struct {
	vertexes  []*graph.Vertex
	edges     []*graph.Edge
	cycleList []*graph.Vertex
}

func NewBellmanFordAlgorithm(vertexes []*graph.Vertex, edges []*graph.Edge) BellmanFordAlgorithm {
	return BellmanFordAlgorithm{vertexes: vertexes, edges: edges, cycleList: make([]*graph.Vertex, 0)}
}

func (this *BellmanFordAlgorithm) Run(source *graph.Vertex) {
	source.MinDistance = 0

	for i := 0; i < len(this.vertexes)-1; i++ {

		for _, e := range this.edges {
			u := e.GetStartVertex()
			v := e.GetTarget()
			distanceCandidate := u.MinDistance + e.GetWeight()

			if distanceCandidate < v.MinDistance {
				v.Predecessor = u
				v.MinDistance = distanceCandidate
			}
		}
	}

	for _, e := range this.edges {
		if e.GetStartVertex().MinDistance != math.MaxFloat64 {
			if hasCycle(e) {
				fmt.Println("There is a negative cycle in", e)
				vertex := e.GetStartVertex()
				for vertex.GetName() != e.GetTarget().GetName() {
					this.cycleList = append(this.cycleList, vertex)
					vertex = vertex.Predecessor
				}
				this.cycleList = append(this.cycleList, e.GetTarget())
				return
			}
		}
	}
}

func hasCycle(edge *graph.Edge) bool {
	return edge.GetStartVertex().MinDistance+edge.GetWeight() < edge.GetTarget().MinDistance
}

func (this *BellmanFordAlgorithm) PrintCycled() {
	for _, v := range this.cycleList {
		fmt.Println(v.GetName(), v.MinDistance)
	}
}
