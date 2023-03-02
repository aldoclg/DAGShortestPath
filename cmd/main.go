package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/aldoclg/DAGShortestPath/dag"
	"github.com/aldoclg/DAGShortestPath/models/graph"
)

var wg sync.WaitGroup

func main() {

	graphList := createGraphList()
	shortestPath := dag.NewShortestPath(graphList)
	shortestPath.Compute()

	for _, v := range graphList {
		fmt.Printf("Distance %.1f, shortest route: %v\n", v.MinDistance, *v)
	}

	dg := dag.DijkstraAlgorithm{}

	graph := createGraphList2()

	graph3, edges3 := createGraphList3()

	bellmanFordAlgorithm := dag.NewBellmanFordAlgorithm(graph3, edges3)

	wg.Add(2)

	go func() {
		defer wg.Done()
		t1 := time.Now()
		bellmanFordAlgorithm.Run(graph3[0])
		fmt.Printf("[BellmanFordAlgorithm] %v %.1f\n", graph3[3], graph3[3].MinDistance)
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))

	}()

	go func() {
		defer wg.Done()
		t1 := time.Now()
		dg.Compute(graph[0])
		fmt.Printf("[DijkstraAlgorithm] %v %.1f\n", graph[3], graph[3].MinDistance)
		t2 := time.Now()
		fmt.Println(t2.Sub(t1))
	}()

	wg.Wait()

	graph4, edges4 := createGraphList4()

	bellmanFordAlgorithm = dag.NewBellmanFordAlgorithm(graph4, edges4)

	bellmanFordAlgorithm.Run(graph4[0])

	fmt.Printf("[FOREX] %v %.3f\n", graph3[3], graph3[3].MinDistance)

	bellmanFordAlgorithm.PrintCycled()

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

func createGraphList3() ([]*graph.Vertex, []*graph.Edge) {
	graphList := make([]*graph.Vertex, 8)

	edgeList := make([]*graph.Edge, 0)

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

	for _, v := range graphList {
		for _, e := range v.GetAdjacency() {
			edgeList = append(edgeList, e)
		}
	}

	return graphList, edgeList
}

func createGraphList4() ([]*graph.Vertex, []*graph.Edge) {

	graphList := make([]*graph.Vertex, 5)

	v0 := graph.NewVertex("USD")
	v1 := graph.NewVertex("EUR")
	v2 := graph.NewVertex("GBP")
	v3 := graph.NewVertex("CHF")
	v4 := graph.NewVertex("CAD")

	graphList[0] = v0
	graphList[1] = v1
	graphList[2] = v2
	graphList[3] = v3
	graphList[4] = v4

	edgeList := make([]*graph.Edge, 0)

	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v0, v1, -1*math.Log(0.741)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v0, v2, -1*math.Log(0.657)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v0, v3, -1*math.Log(1.061)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v0, v4, -1*math.Log(1.005)))

	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v1, v0, -1*math.Log(1.349)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v1, v2, -1*math.Log(0.888)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v1, v3, -1*math.Log(1.433)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v1, v4, -1*math.Log(1.366)))

	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v2, v0, -1*math.Log(1.521)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v2, v1, -1*math.Log(1.126)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v2, v3, -1*math.Log(1.614)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v2, v4, -1*math.Log(1.538)))

	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v3, v0, -1*math.Log(0.942)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v3, v1, -1*math.Log(0.698)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v3, v2, -1*math.Log(0.619)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v3, v4, -1*math.Log(0.953)))

	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v4, v0, -1*math.Log(0.995)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v4, v1, -1*math.Log(0.732)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v4, v2, -1*math.Log(0.650)))
	edgeList = append(edgeList, graph.NewEdgeWithStartVertex(v4, v3, -1*math.Log(1.049)))

	return graphList, edgeList

}
