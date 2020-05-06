package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}

/*
	a - b - d
	|	    |
	c - e	f
*/
func main() {
	addEdge("a", "b")
	addEdge("a", "c")
	addEdge("b", "d")
	addEdge("c", "e")
	addEdge("d", "f")
	fmt.Println("a to c = ", hasEdge("a", "c"))
	fmt.Println("a to b = ", hasEdge("a", "b"))
	fmt.Println("b to d = ", hasEdge("b", "d"))

	fmt.Println("a to d = ", hasEdge("a", "d"))
	fmt.Println("a to f = ", hasEdge("a", "f"))

}
