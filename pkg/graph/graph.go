package graph

import "fmt"

type AdjVertex struct {
	I int
	Next *AdjVertex
}

type AdjList struct {
	Head *AdjVertex
}

type Graph struct {
	graph map[int]*AdjList
	V int
}

// Undirected Graph - add edge
func (g *Graph) AddEdge(u, v int) {
	vertexU := AdjVertex{I: u, Next: nil}
	vertexV := AdjVertex{I: v, Next: nil}

	adjListU := (*g).graph[u]
	adjListV := (*g).graph[v]

	// add edge : u -> v
	addEdge(g, adjListU, vertexV, u)
	// add edge : v -> u
	addEdge(g, adjListV, vertexU, v)
}

// Directed Graph - add edge
func (g *Graph) AddDirEdge(u, v int) {
	vertexV := AdjVertex{I: v, Next: nil}

	adjListU := (*g).graph[u]

	// add edge : u -> v
	addEdge(g, adjListU, vertexV, u)
}

//private func to add edge
func addEdge(g *Graph, adjList *AdjList, vertex AdjVertex, t int) {

	// Check if adj list already exists
	if adjList != nil {
		//take a head pointer
		head := adjList.Head
		for head.Next != nil {
			head = head.Next
		}
		head.Next = &vertex
	} else {
		adjL := AdjList{Head: &vertex}
		(*g).graph[t] = &adjL
	}

}

func (g *Graph) Bfs(s int) {

	var queue []int
	visited := make([]bool, g.V)

	visited[s] = true
	queue = append(queue, s)

	var head *AdjVertex
	for len(queue) != 0 {
		e := queue[0]
		fmt.Println(e)
		//Pop the first element and update the queue
		queue = queue[1:]

		if (*g).graph[e] != nil {
			head = (*g).graph[e].Head
		}

		for head != nil {
			if !visited[head.I] {
				queue = append(queue, head.I)
				visited[head.I] = true
			}
			head = head.Next
		}
	}

}

// Driver program
func main() {
	 //Directed Graph
	dg := Graph{V:6, graph: make(map[int]*AdjList)}
	dg.AddDirEdge(2, 1)
	dg.AddDirEdge(0, 2)
	dg.AddDirEdge(0, 3)
	dg.AddDirEdge(0, 4)
	dg.AddDirEdge(0, 5)
	dg.AddDirEdge(1, 3)
	dg.AddDirEdge(2, 3)
	fmt.Println("* * * * * BFS for Directed Graph * * * * ")
	dg.Bfs(0)

	g := Graph{V:6, graph: make(map[int]*AdjList)}
	g.AddEdge(2, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(0, 5)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	fmt.Println("* * * * * BFS for Undirected Graph * * * * ")
	g.Bfs(0)

	// Print Graph
	fmt.Println("* * * * * Print Graph * * * * ")
	for i := 0; i < len(g.graph) ; i++ {
		head := g.graph[i].Head
		fmt.Print(i, "--")
		for head != nil {
			fmt.Print("->", head.I)
			head = head.Next
		}
		fmt.Println("")
	}

}

