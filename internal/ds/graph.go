package ds

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

// Graph
type Graph struct {
	AllEdges    []*Edge
	AllVertices map[int]*Vertex
}

func (g *Graph) AddVertex(vertex *Vertex) {

	if _, ok := g.AllVertices[(*vertex).Id]; ok {
		return
	}
	g.AllVertices[(*vertex).Id] = vertex
	for _, edgeptr := range (*vertex).Edges {
		g.AllEdges = append(g.AllEdges, edgeptr)
	}

}

func (g *Graph) GetVertex(id int) *Vertex {
	return g.AllVertices[id]
}

func (g *Graph) AddEdge(id1 int, id2 int, weight int) {

	var v1 *Vertex
	if vertex, ok := g.AllVertices[id1]; ok {
		v1 = vertex
	} else {
		v1 = &Vertex{id1, Coordinates{0, 0}, make([]*Edge, 0), make([]*Vertex, 0)}
		g.AllVertices[id1] = v1
	}

	var v2 *Vertex
	if vertex, ok := g.AllVertices[id2]; ok {
		v2 = vertex
	} else {
		v2 = &Vertex{id2, Coordinates{0, 0}, make([]*Edge, 0), make([]*Vertex, 0)}
		g.AllVertices[id2] = v2
	}

	edge := &Edge{V1: v1, V2: v2, Weight: weight}

	g.AllEdges = append(g.AllEdges, edge)
	v1.AddAdjacentVertex(edge, v2)
	v2.AddAdjacentVertex(edge, v1)

}

func (g *Graph) AddEdge1(edge *Edge) {

	var v1 *Vertex
	id1 := (*edge).V1.Id
	data1 := (*edge).V1.Data
	if vertex, ok := g.AllVertices[id1]; ok {
		v1 = vertex
	} else {
		v1 = &Vertex{id1, data1, make([]*Edge, 0), make([]*Vertex, 0)}
		g.AllVertices[id1] = v1
	}

	var v2 *Vertex
	id2 := (*edge).V2.Id
	data2 := (*edge).V2.Data
	if vertex, ok := g.AllVertices[id2]; ok {
		v2 = vertex
	} else {
		v2 = &Vertex{id2, data2, make([]*Edge, 0), make([]*Vertex, 0)}
		g.AllVertices[id2] = v2
	}

	g.AllEdges = append(g.AllEdges, edge)
	v1.AddAdjacentVertex(edge, v2)
	v2.AddAdjacentVertex(edge, v1)

}

func (g *Graph) GetAllVertices() []*Vertex {

	vertices := make([]*Vertex, 0)

	for _, value := range g.AllVertices {
		vertices = append(vertices, value)
	}

	return vertices
}

func (g *Graph) setDataForVertex(id int, data Coordinates) {

	if v, ok := g.AllVertices[id]; ok {
		v.Data = data
	}

}

func (g *Graph) String() string {
	var out bytes.Buffer

	for _, eptr := range g.AllEdges {
		out.WriteString(eptr.String())
	}

	return out.String()
}

func containsData(id int, nodeMap map[int]int) (int, error) {
	if value, ok := nodeMap[id]; ok {
		return value, nil
	} else {
		return 0, errors.New("no key found")
	}

}

// Edge
type Edge struct {
	V1, V2 *Vertex
	Weight int
}

func (e *Edge) String() string {
	s := strings.Join([]string{"Edge[", "V1 =", (e.V1).String(), "V2 =", (e.V2).String(), "weight =", strconv.Itoa(e.Weight), "]"}, " ")
	return s
}

// Coordinates
type Coordinates struct {
	Row, Col int
}

// Vertex
type Vertex struct {
	Id               int
	Data             Coordinates
	Edges            []*Edge
	AdjacentVertices []*Vertex
}

func (v *Vertex) AddAdjacentVertex(e *Edge, adjacentVertex *Vertex) {

	v.Edges = append(v.Edges, e)
	v.AdjacentVertices = append(v.AdjacentVertices, v)
}

func (v *Vertex) getDegree() int {
	return len(v.Edges)
}

func (v *Vertex) String() string {
	return strconv.Itoa(v.Id)
}
