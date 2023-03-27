package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"mazerunner/internal/algorithms"
	"mazerunner/internal/ds"
	"strings"
)

type PathElement struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func getMSTNew(size int) []*(ds.Edge) {

	graph := &(ds.Graph{AllEdges: make([]*(ds.Edge), 0), AllVertices: make(map[int]*(ds.Vertex))})

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j < size-1 {
				graph.AddEdge(size*i+j, size*i+(j+1), rand.Intn(10))
			}
			if i < size-1 {
				graph.AddEdge(size*i+j, size*(i+1)+j, rand.Intn(10))
			}
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			graph.GetVertex(size*i + j).Data = ds.Coordinates{Row: i, Col: j}
		}
	}

	return algorithms.PrimMST(graph)
}

func solveMaze(maze [][]int, row_start, col_start, row_end, col_end int) ([]PathElement, error) {

	Size := len(maze)
	actualMazeSize := (Size + 1) / 2

	start := actualMazeSize*row_start + col_start
	end := actualMazeSize*row_end + col_end

	var i, j int

	mst := reconstructMSTfromMatrix(maze)

	dsp := algorithms.ShortestPath(mst, mst.GetVertex(start))

	//jprev := 2 * (mst.GetVertex(end).Data.Col)
	//iprev := 2 * (mst.GetVertex(end).Data.Row)
	//maze[iprev][jprev] = 2

	if _, ok := dsp[end]; !ok {
		err := errors.New("no path found")
		return nil, err
	}

	path := make([]PathElement, 0)

	k, ok := dsp[end]

	for k != -1 && ok {
		i = 2 * (mst.GetVertex(k).Data.Row)
		j = 2 * (mst.GetVertex(k).Data.Col)
		//maze[i][j] = 2
		//maze[(iprev+i)/2][(jprev+j)/2] = 2
		//iprev, jprev = i, j
		path = append(path, PathElement{Row: i, Col: j})
		k, ok = dsp[k]

	}

	return path, nil
}

func generateMaze(size int, edges []*(ds.Edge)) [][]int {

	Size := 2*size - 1

	var maze [][]int

	for row := 0; row < Size; row++ {

		sl := make([]int, 0)
		for col := 0; col < Size; col++ {

			if row%2 == 0 && col%2 == 0 {
				sl = append(sl, 1)
			} else {
				sl = append(sl, 0)
			}
		}
		maze = append(maze, sl)
	}

	for _, edge := range edges {

		v1 := (*edge).V1
		v2 := (*edge).V2

		rowNumber := (*v1).Data.Row + (*v2).Data.Row
		colNumber := (*v1).Data.Col + (*v2).Data.Col

		maze[rowNumber][colNumber] = 1
	}

	return maze
}

func genMazeString(maze [][]int) string {

	var sb strings.Builder

	for row := 0; row < len(maze); row++ {

		for col := 0; col < len(maze); col++ {

			fmt.Fprintf(&sb, "%d", maze[row][col])
		}
	}

	return sb.String()
}

func reconstructMatrix(mazeString string) [][]int {

	maze := make([][]int, 0)
	mazeSize := int(math.Sqrt(float64(len(mazeString))))

	for i := 0; i < mazeSize; i++ {

		sl := make([]int, 0)

		for j := 0; j < mazeSize; j++ {
			sl = append(sl, int(mazeString[i*mazeSize+j])-48)
		}

		maze = append(maze, sl)
	}
	return maze
}

func reconstructMSTfromMatrix(maze [][]int) *(ds.Graph) {

	Size := len(maze)
	actualMazeSize := (Size + 1) / 2

	mstGraph := &(ds.Graph{
		AllEdges:    make([]*(ds.Edge), 0),
		AllVertices: make(map[int]*(ds.Vertex)),
	})

	for i := 0; i < actualMazeSize; i++ {

		for j := 0; j < actualMazeSize; j++ {

			if j < actualMazeSize-1 {

				if maze[2*i][2*j+1] == 1 {

					mstGraph.AddEdge(actualMazeSize*i+j, actualMazeSize*i+(j+1), 1)
				}
			}

			if i < actualMazeSize-1 {

				if maze[2*i+1][2*j] == 1 {

					mstGraph.AddEdge(actualMazeSize*i+j, actualMazeSize*(i+1)+j, 1)
				}
			}
		}
	}

	for i := 0; i < actualMazeSize; i++ {

		for j := 0; j < actualMazeSize; j++ {

			mstGraph.GetVertex(actualMazeSize*i + j).Data = ds.Coordinates{Row: i, Col: j}
		}
	}

	return mstGraph
}

type envelope map[string]interface{}

func writeJSON(data envelope) ([]byte, error) {

	js, err := json.Marshal(data)
	if err != nil {

		return nil, err
	}
	js = append(js, '\n')

	return js, nil
}

func readJSON(input []byte, dst interface{}) error {

	maxBytes := 1_048_576
	dec := json.NewDecoder(bytes.NewReader(input))
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {

		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {

		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.As(err, &unmarshalTypeError):

			if unmarshalTypeError.Field != "" {

				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}

			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:

			return err

		}

	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}
