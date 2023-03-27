package algorithms

import (
	"errors"
	"fmt"
)

type node struct {
	weight   int
	vertexId int
}

type BinaryMinHeap struct {
	nodePtrs     []*node
	nodePosition map[int]int
}

func NewMinHeap() *BinaryMinHeap {
	return &BinaryMinHeap{make([]*node, 0), make(map[int]int)}
}

func (binaryMinHeap *BinaryMinHeap) ContainsData(vertexId int) bool {

	_, ok := binaryMinHeap.nodePosition[vertexId]
	return ok
}

func (binaryMinHeap *BinaryMinHeap) Add(weight int, vertexId int) {

	n := &node{weight, vertexId}
	binaryMinHeap.nodePtrs = append(binaryMinHeap.nodePtrs, n)
	size := len(binaryMinHeap.nodePtrs)
	current := size - 1
	parentIndex := (current - 1) / 2
	binaryMinHeap.nodePosition[vertexId] = current

	for parentIndex >= 0 {
		parentNode := binaryMinHeap.nodePtrs[parentIndex]
		currentNode := binaryMinHeap.nodePtrs[current]

		if (*parentNode).weight > (*currentNode).weight {
			binaryMinHeap.swap(parentNode, currentNode)
			binaryMinHeap.updatePositionMap(parentNode, currentNode, parentIndex, current)

			current = parentIndex
			parentIndex = (parentIndex - 1) / 2

			currentNode = binaryMinHeap.nodePtrs[current]
			parentNode = binaryMinHeap.nodePtrs[parentIndex]
		} else {
			break
		}
	}

}

func (binaryMinHeap *BinaryMinHeap) min() int {
	return (*(binaryMinHeap.nodePtrs[0])).vertexId
}

func (binaryMinHeap *BinaryMinHeap) Empty() bool {
	return len(binaryMinHeap.nodePtrs) == 0
}

func (binaryMinHeap *BinaryMinHeap) Decrease(vertexId int, newWeight int) {

	if position, ok := binaryMinHeap.nodePosition[vertexId]; ok {
		currentNode := binaryMinHeap.nodePtrs[position]
		(*currentNode).weight = newWeight
		parentIndex := (position - 1) / 2
		parentNode := binaryMinHeap.nodePtrs[parentIndex]

		for parentIndex >= 0 {
			if (*parentNode).weight > (*currentNode).weight {

				binaryMinHeap.swap(parentNode, currentNode)
				binaryMinHeap.updatePositionMap(parentNode, currentNode, parentIndex, position)
				position = parentIndex
				parentIndex = (parentIndex - 1) / 2

				currentNode = binaryMinHeap.nodePtrs[position]
				parentNode = binaryMinHeap.nodePtrs[parentIndex]

			} else {
				break
			}
		}
	}

}

func (binaryMinHeap *BinaryMinHeap) ExtractMinNode() *node {

	size := len(binaryMinHeap.nodePtrs) - 1

	n1 := binaryMinHeap.nodePtrs[0]
	n2 := binaryMinHeap.nodePtrs[size]

	minNode := &node{(*n1).weight, (*n1).vertexId}
	(*n1).weight = (*n2).weight
	(*n1).vertexId = (*n2).vertexId

	delete(binaryMinHeap.nodePosition, (*minNode).vertexId)
	binaryMinHeap.nodePosition[(*n2).vertexId] = 0

	binaryMinHeap.nodePtrs = binaryMinHeap.nodePtrs[:size]

	currentIndex := 0
	size--

	for {
		left := 2*currentIndex + 1
		right := 2*currentIndex + 2

		if left > size {
			break
		}
		if right > size {
			right = left
		}

		var smallerIndex int

		if (*(binaryMinHeap.nodePtrs[left])).weight <= (*(binaryMinHeap.nodePtrs[right])).weight {
			smallerIndex = left
		} else {
			smallerIndex = right
		}

		if (*(binaryMinHeap.nodePtrs[currentIndex])).weight > (*(binaryMinHeap.nodePtrs[smallerIndex])).weight {

			binaryMinHeap.swap(binaryMinHeap.nodePtrs[currentIndex], binaryMinHeap.nodePtrs[smallerIndex])
			binaryMinHeap.updatePositionMap(binaryMinHeap.nodePtrs[currentIndex], binaryMinHeap.nodePtrs[smallerIndex], currentIndex, smallerIndex)
			currentIndex = smallerIndex
		} else {
			break
		}

	}

	return minNode
}

func (binaryMinHeap *BinaryMinHeap) ExtractMin() int {

	n := binaryMinHeap.ExtractMinNode()
	return (*n).vertexId
}

func (binaryMinHeap *BinaryMinHeap) GetWeight(vertexId int) (int, error) {

	if position, ok := binaryMinHeap.nodePosition[vertexId]; ok {
		return (*(binaryMinHeap.nodePtrs[position])).weight, nil
	} else {
		return 0, errors.New("No such key exists")
	}

}

func (binaryMinHeap *BinaryMinHeap) swap(node1 *node, node2 *node) {

	vertexId := (*node2).vertexId
	weight := (*node2).weight

	(*node2).vertexId = (*node1).vertexId
	(*node2).weight = (*node1).weight

	(*node1).vertexId = vertexId
	(*node1).weight = weight

}

func (binaryMinHeap *BinaryMinHeap) updatePositionMap(node1 *node, node2 *node, node1Index int, node2Index int) {

	binaryMinHeap.nodePosition[(*node1).vertexId] = node1Index
	binaryMinHeap.nodePosition[(*node2).vertexId] = node2Index

}

func (binaryMinHeap *BinaryMinHeap) printHeap() {
	for _, nodePtr := range binaryMinHeap.nodePtrs {
		fmt.Println("vertexId:", (*nodePtr).vertexId, "weight", (*nodePtr).weight)
	}
}

func (binaryMinHeap *BinaryMinHeap) printPositionMap() {
	fmt.Println(binaryMinHeap.nodePosition)
}
