package main

import "fmt"

func dfs(node int, visited []bool, graph map[int][]int) {

	visited[node] = true

	fmt.Printf("%d ", node)

	for _, children := range graph[node] {
		if !visited[children] {
			dfs(children, visited, graph)
		}

	}

}

func main() {

	graph := map[int][]int{

		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 4},
		3: {1, 5},
		4: {1, 2},
		5: {3},
	}

	visited := make([]bool, len(graph)+1)

	dfs(0, visited, graph)

}
