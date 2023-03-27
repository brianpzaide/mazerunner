package algorithms

import (
	"math"
	"mazerunner/internal/ds"
)

func ShortestPath(graph *(ds.Graph), sourceVertex *(ds.Vertex)) map[int]int {

	minHeap := NewMinHeap()
	distanceMap := make(map[int]int)
	parentMap := make(map[int]int)

	for _, v := range (*graph).GetAllVertices() {
		minHeap.Add(math.MaxInt64, (*v).Id)
	}
	//fmt.Println("printing minHeap:")
	//fmt.Println((*sourceVertex).Id)

	minHeap.Decrease((*sourceVertex).Id, 0)

	distanceMap[(*sourceVertex).Id] = 0

	parentMap[(*sourceVertex).Id] = -1

	for !minHeap.Empty() {

		n := minHeap.ExtractMinNode()
		currentVertexId := (*n).vertexId
		currentWeight := (*n).weight
		current := graph.GetVertex(currentVertexId)
		distanceMap[currentVertexId] = currentWeight

		for _, edge := range (*current).Edges {
			//fmt.Println(current)
			//fmt.Println(edge.V1, edge.V2)
			adjacent := getVertexForEdge(current, edge)
			//fmt.Println(adjacent)
			//fmt.Println("==============================")
			if minHeap.ContainsData((*adjacent).Id) {
				newDistance := distanceMap[currentVertexId] + (*edge).Weight
				w, err := minHeap.GetWeight((*adjacent).Id)
				if err == nil && w > newDistance {
					minHeap.Decrease((*adjacent).Id, newDistance)
					parentMap[(*adjacent).Id] = currentVertexId
				}
			}

		}
	}
	//fmt.Println("==============================")
	return parentMap
}
