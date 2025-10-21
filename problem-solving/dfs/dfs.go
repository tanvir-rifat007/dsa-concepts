package main

import "fmt"

func dfsIterative(start int, graph map[int][]int) {

	visited := make(map[int]bool)
	stack := []int{start}

	visited[start] = true

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("%v ", node)

		for i := len(graph[node]) - 1; i >= 0; i-- {

			child := graph[node][i]
			if !visited[child] {
				visited[child] = true
				stack = append(stack, child)
			}
		}

	}

}

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

	//visited := make([]bool, len(graph)+1)

	dfsIterative(0, graph)

}
