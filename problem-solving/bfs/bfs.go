package main

import "fmt"

func BFS(start int, graph map[int][]int) {

	visited := make(map[int]bool)

	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node)

		for _, children := range graph[node] {

			if !visited[children] {
				visited[children] = true
				queue = append(queue, children)
			}

		}

	}

}

func main() {

	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {},
	}

	fmt.Print("BFS traversal: ")
	BFS(1, graph)
}
