# mazerunner

A simple app made using HTML5 canvas and go-wasm

Prim's mst algorithm is used to generate a random square maze.
Dijkstra's shortest path algorithm is used to find a shortest path between any two points chosen by the user.

#### To build
copy the javascript glue code that already comes with the go installation into the ```out``` directory

```cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./out/```

compile the go program to webassembly
```
cd app
GOOS=js GOARCH=wasm go build -o ../out/mr.wasm
```
#### To run
change into the ```out``` directory and run the python http server

```python3 -m http.server```

point the browser to ```localhost:8000``` 

[Demo](https://brianpzaide.github.io/mazerunner)
