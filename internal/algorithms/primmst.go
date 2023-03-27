package algorithms

import (
	"math"
	"mazerunner/internal/ds"
)

func PrimMST(graph *(ds.Graph)) []*(ds.Edge) {

	minHeap := NewMinHeap()
	vertexToEdge := make(map[int]*(ds.Edge))
	result := make([]*(ds.Edge), 0)

	for _, v := range (*graph).GetAllVertices() {
		minHeap.Add(math.MaxInt64, (*v).Id)
	}

	startVertex := (*graph).GetAllVertices()[0]

	minHeap.Decrease((*startVertex).Id, 0)

	for !minHeap.Empty() {

		current := (*graph).GetVertex(minHeap.ExtractMin())
		if spanningTreeEdge, ok := vertexToEdge[(*current).Id]; ok {
			result = append(result, spanningTreeEdge)
		}

		for _, e := range (*current).Edges {

			adjacent := getVertexForEdge(current, e)

			if minHeap.ContainsData((*adjacent).Id) {
				w, err := minHeap.GetWeight((*adjacent).Id)
				if err == nil && w > (*e).Weight {
					minHeap.Decrease((*adjacent).Id, (*e).Weight)
					vertexToEdge[(*adjacent).Id] = e
				}
			}

		}
	}
	return result
}

func getVertexForEdge(v *(ds.Vertex), e *(ds.Edge)) *(ds.Vertex) {

	//fmt.Println(e.V1, e.V2)

	if (*((*e).V1)).Id == (*v).Id {
		return (*e).V2
	} else {
		return (*e).V1
	}

}
