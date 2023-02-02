package main

import "fmt"

type Graph struct {
	adjacencyList map[*Vertex][]*Vertex
	isDirected    bool
}

func NewGraph(isDirected bool) Graph {
	adjacencyList := make(map[*Vertex][]*Vertex)
	return Graph{adjacencyList: adjacencyList, isDirected: isDirected}
}

type Vertex struct {
	Data interface{}
}

func (g Graph) AddVertex(data interface{}) *Vertex {
	v := Vertex{
		Data: data,
	}

	g.adjacencyList[&v] = make([]*Vertex, 0)

	return &v
}

func (g Graph) AddEdge(src, dst *Vertex) {
	g.adjacencyList[src] = append(g.adjacencyList[src], dst)
	if g.isDirected {
		return
	}

	g.adjacencyList[dst] = append(g.adjacencyList[dst], src)
}

func (g Graph) Print() {
	for vertex, neighbours := range g.adjacencyList {
		fmt.Print(vertex.Data)
		fmt.Print("->")
		for _, n := range neighbours {
			fmt.Print(n.Data)
			fmt.Print(",")
		}
		fmt.Println()
	}
}

func (g Graph) Bfs(start *Vertex) []interface{} {
	visited := make(map[*Vertex]bool)
	queue := make([]*Vertex, 0)
	values := make([]interface{}, 0)

	queue = append(queue, start)

	for len(queue) != 0 {
		// pop
		top := queue[0]
		// trim
		queue = queue[1:]

		// already visited - saves from loop.
		if visited[top] {
			continue
		}

		//mark popped node visited.
		visited[top] = true
		// collect the value.
		values = append(values, top.Data)
		// queue the neighbours.
		queue = append(queue, g.adjacencyList[top]...)
	}

	return values
}

func (g Graph) Dfs(start *Vertex) []interface{} {
	visited := make(map[*Vertex]bool)
	values := make([]interface{}, 0)

	g.dfs(start, visited, &values)

	return values
}

func (g Graph) dfs(start *Vertex, visited map[*Vertex]bool, values *[]interface{}) {
	if visited[start] {
		return
	}
	visited[start] = true
	*values = append(*values, start.Data)

	for _, neighbour := range g.adjacencyList[start] {
		g.dfs(neighbour, visited, values)
	}
}

func main() {
	g := NewGraph(true)

	v1 := g.AddVertex(1)
	v2 := g.AddVertex(2)
	v3 := g.AddVertex(3)
	v4 := g.AddVertex(4)
	v5 := g.AddVertex(5)
	v6 := g.AddVertex(6)

	g.AddEdge(v1, v2)
	g.AddEdge(v2, v3)
	g.AddEdge(v1, v4)
	g.AddEdge(v3, v4)
	g.AddEdge(v4, v4)
	g.AddEdge(v3, v5)
	g.AddEdge(v3, v6)

	g.Print()
	fmt.Println("---- BFS ----")
	for _, d := range g.Bfs(v1) {
		fmt.Print(d)
		fmt.Print(",")
	}

	fmt.Println("\n---- DFS ----")
	for _, d := range g.Dfs(v1) {
		fmt.Print(d)
		fmt.Print(",")
	}
}
