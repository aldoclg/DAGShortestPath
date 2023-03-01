package main

import (
	"fmt"

	"github.com/aldoclg/DAGShortestPath/dag"
	"github.com/aldoclg/DAGShortestPath/models/graph"
)

func main() {

	graphList := createGraphList()
	shortestPath := dag.NewShortestPath(graphList)
	shortestPath.Compute()

	for _, v := range graphList {
		fmt.Printf("Distance %d, shortest route: %v\n", v.MinDistance, *v)
	}

	dg := dag.DijkstraAlgorithm{}

	graph := createGraphList2()

	dg.Compute(graph[0])
	fmt.Printf("%v %d\n", graph[6], graph[6].MinDistance)

}

func createGraphList() []*graph.Vertex {
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

	return graphList
}

func createGraphList2() []*graph.Vertex {
	graphList := make([]*graph.Vertex, 8)

	v0 := graph.NewVertex("A")
	v1 := graph.NewVertex("B")
	v2 := graph.NewVertex("C")
	v3 := graph.NewVertex("D")
	v4 := graph.NewVertex("E")
	v5 := graph.NewVertex("F")
	v6 := graph.NewVertex("G")
	v7 := graph.NewVertex("H")

	v0.AddNeighbor(graph.NewEdgeWithStartVertex(v0, v1, 5))
	v0.AddNeighbor(graph.NewEdgeWithStartVertex(v0, v4, 9))
	v0.AddNeighbor(graph.NewEdgeWithStartVertex(v0, v7, 8))

	v1.AddNeighbor(graph.NewEdgeWithStartVertex(v1, v2, 12))
	v1.AddNeighbor(graph.NewEdgeWithStartVertex(v1, v3, 15))
	v1.AddNeighbor(graph.NewEdgeWithStartVertex(v1, v7, 6))

	v2.AddNeighbor(graph.NewEdgeWithStartVertex(v2, v3, 3))
	v2.AddNeighbor(graph.NewEdgeWithStartVertex(v2, v6, 11))

	v3.AddNeighbor(graph.NewEdgeWithStartVertex(v3, v6, 9))

	v4.AddNeighbor(graph.NewEdgeWithStartVertex(v4, v7, 5))
	v4.AddNeighbor(graph.NewEdgeWithStartVertex(v4, v5, 4)) //13
	v4.AddNeighbor(graph.NewEdgeWithStartVertex(v4, v6, 20))

	v5.AddNeighbor(graph.NewEdgeWithStartVertex(v5, v6, 13))
	v5.AddNeighbor(graph.NewEdgeWithStartVertex(v5, v2, 1)) //14

	v7.AddNeighbor(graph.NewEdgeWithStartVertex(v7, v2, 7))
	v7.AddNeighbor(graph.NewEdgeWithStartVertex(v7, v5, 6))

	graphList[0] = v0
	graphList[1] = v1
	graphList[2] = v2
	graphList[3] = v3
	graphList[4] = v4
	graphList[5] = v5
	graphList[6] = v6
	graphList[7] = v7

	return graphList
}
