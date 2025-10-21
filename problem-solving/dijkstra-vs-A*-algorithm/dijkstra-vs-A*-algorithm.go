package main

import (
	"container/heap"
	"fmt"
	"math"
)

// -------------------- Priority Queue --------------------
type Item struct {
	node     string
	priority float64
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// -------------------- Graph Definition --------------------
type Graph map[string]map[string]float64

func makeGraph() Graph {
	return Graph{
		"A": {"B": 2, "C": 1},
		"B": {"D": 5},
		"C": {"D": 2, "E": 1},
		"D": {"E": 3},
		"E": {},
	}
}

// -------------------- Dijkstra --------------------
func Dijkstra(g Graph, start, goal string) (float64, []string, int) {
	dist := make(map[string]float64)
	prev := make(map[string]string)
	for node := range g {
		dist[node] = math.Inf(1)
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, priority: 0})

	steps := 0 // count how many nodes are explored

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node := item.node
		steps++

		if node == goal {
			break
		}

		for neighbor, weight := range g[node] {
			newDist := dist[node] + weight
			if newDist < dist[neighbor] {
				dist[neighbor] = newDist
				prev[neighbor] = node
				heap.Push(pq, &Item{node: neighbor, priority: newDist})
			}
		}
	}

	return dist[goal], reconstructPath(prev, start, goal), steps
}

// -------------------- A* --------------------
func AStar(g Graph, start, goal string, heuristic map[string]float64) (float64, []string, int) {
	dist := make(map[string]float64)
	prev := make(map[string]string)
	for node := range g {
		dist[node] = math.Inf(1)
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, priority: heuristic[start]})

	steps := 0 // count how many nodes are explored

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node := item.node
		steps++

		if node == goal {
			break
		}

		for neighbor, weight := range g[node] {
			newDist := dist[node] + weight
			if newDist < dist[neighbor] {
				dist[neighbor] = newDist
				prev[neighbor] = node
				f := newDist + heuristic[neighbor]
				heap.Push(pq, &Item{node: neighbor, priority: f})
			}
		}
	}

	return dist[goal], reconstructPath(prev, start, goal), steps
}

// -------------------- Helper --------------------
func reconstructPath(prev map[string]string, start, goal string) []string {
	path := []string{}
	for at := goal; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
		if at == start {
			break
		}
	}
	return path
}

// -------------------- Main --------------------
func main() {
	graph := makeGraph()

	heuristic := map[string]float64{
		"A": 3,
		"B": 2,
		"C": 1,
		"D": 0,
		"E": 1,
	}

	// Dijkstra
	costD, pathD, stepsD := Dijkstra(graph, "A", "D")
	fmt.Printf("Dijkstra: cost = %.0f, path = %v, steps = %d\n", costD, pathD, stepsD)

	// A*
	costA, pathA, stepsA := AStar(graph, "A", "D", heuristic)
	fmt.Printf("A*: cost = %.0f, path = %v, steps = %d\n", costA, pathA, stepsA)
}
