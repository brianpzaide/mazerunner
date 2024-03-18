# mazerunner

Mazerunner is an application that generates random square mazes of varyinge sizes and finds the shortest path between any two points within the maze. It utilizes HTML5 canvas for rendering and Go WebAssembly (go-wasm) for backend processing.
Prim's minimal spanning tree algorithm is used to generate random square mazes, while Dijkstra's shortest path algorithm is used to calculate the shortest path between user-selected points.

### Features

- Generate random square mazes of varying sizes.
- Find the shortest path between any two points within the maze.
- Interactive visualization using HTML5 canvas.

### To build
copy the javascript glue code that already comes with the go installation into the ```out``` directory

```cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./out/```

compile the go program to webassembly
```
cd app
GOOS=js GOARCH=wasm go build -o ../out/mr.wasm
```
### To run
change to the ```out``` directory and start the python HTTP server

```python3 -m http.server```

Open your web browser and navigate to ```localhost:8000```. Interact with the maze generator and pathfinding features directly in your browser.  

### Demo
Check out the [Demo](https://brianpzaide.github.io/mazerunner) to see Mazerunner in action.


