package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("createMaze", createWrapper())
	js.Global().Set("solveMaze", solveWrapper())
	<-make(chan bool)
}

func create(size int) (string, error) {
	if size < 3 {
		return "", errors.New("size must be an integer greater than 2")
	}

	edges := getMSTNew(size)
	maze := generateMaze(size, edges)
	mazeString := genMazeString(maze)

	resp, err := writeJSON(envelope{"mazeSize": size, "cellSize": 25, "mazeString": mazeString})
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func createWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputSize := args[0].String()
		size, err := strconv.Atoi(inputSize)
		if err != nil {
			return "size must be an integer"
		}

		resp, err := create(size)
		if err != nil {
			return err.Error()
		}
		return resp
	})
	return jsonFunc

}

func solve(inputbytes []byte) (string, error) {

	var input struct {
		MazeString string `json:"mazeString"`
		RowStart   int    `json:"rowStart"`
		ColStart   int    `json:"colStart"`
		RowEnd     int    `json:"rowEnd"`
		ColEnd     int    `json:"colEnd"`
	}

	err := readJSON(inputbytes, &input)
	if err != nil {
		return "", err
	}

	reconstructedMaze := reconstructMatrix(input.MazeString)

	path, err := solveMaze(reconstructedMaze, input.RowStart, input.ColStart, input.RowEnd, input.ColEnd)
	if err != nil {
		return "", err
	}

	resp, err := writeJSON(envelope{"path": path})
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func solveWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputbytes := args[0].String()
		resp, err := solve([]byte(inputbytes))
		if err != nil {
			return err.Error()
		}
		return resp
	})
	return jsonFunc

}
