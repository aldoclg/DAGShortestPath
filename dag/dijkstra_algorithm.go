package dag

import (
	"github.com/aldoclg/DAGShortestPath/models/graph"
	"github.com/aldoclg/DAGShortestPath/models/queue"
)

type DijkstraAlgorithm struct {
}

func (da *DijkstraAlgorithm) Compute(vertex *graph.Vertex) {
	vertex.MinDistance = 0

	priorityQueue := queue.NewPriorityQueue[*graph.Vertex]()
	priorityQueue.Offer(vertex)

	for priorityQueue.IsNotEmpty() {

		actual := priorityQueue.Poll()
		for _, e := range actual.GetAdjacency() {

			t := e.GetTarget()
			distanceCandidate := actual.MinDistance + e.GetWeight()

			if distanceCandidate < t.MinDistance {
				priorityQueue.Remove(t)
				t.MinDistance = distanceCandidate
				t.Predecessor = actual
				priorityQueue.Offer(t)
			}
		}
	}
}
