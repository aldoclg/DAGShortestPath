package main

import (
	"fmt"

	"github.com/aldoclg/DAGShortestPath/dag"
	"github.com/aldoclg/DAGShortestPath/models/graph"
)

func main() {
	graphList := make([]*graph.Vertex, 6)

	v0 := graph.NewVertex("S")
	v1 := graph.NewVertex("A")
	v2 := graph.NewVertex("B")
	v3 := graph.NewVertex("C")
	v4 := graph.NewVertex("D")
	v5 := graph.NewVertex("E")

	v0.AddNeighbor(graph.NewEdge(v1, 1))
	v0.AddNeighbor(graph.NewEdge(v3, 2))

	v1.AddNeighbor(graph.NewEdge(v2, 6))

	v2.AddNeighbor(graph.NewEdge(v4, 1))
	v2.AddNeighbor(graph.NewEdge(v5, 2))

	v3.AddNeighbor(graph.NewEdge(v1, 4))
	v3.AddNeighbor(graph.NewEdge(v4, 3))

	v4.AddNeighbor(graph.NewEdge(v5, 1))

	graphList[0] = v0
	graphList[1] = v1
	graphList[2] = v2
	graphList[3] = v3
	graphList[4] = v4
	graphList[5] = v5

	shortestPath := dag.NewShortestPath(graphList)
	shortestPath.Compute()

	for _, v := range graphList {
		fmt.Printf("Distance %d, shortest route: %v\n", v.MinDistance, *v)
	}

}
