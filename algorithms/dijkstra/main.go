package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

type Item struct {
	node     int
	distance int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(adj [][]Edge, start int) []int {
	n := len(adj)
	dist := make([]int, n)
	parent := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		u := current.node

		if current.distance > dist[u] {
			continue
		}

		for _, edge := range adj[u] {
			v := edge.to
			if dist[u] + edge.weight < dist[v] {
				dist[v] = dist[u] + edge.weight
				parent[v] = u
				heap.Push(&pq, &Item{node: v, distance: dist[v]})
			}
		}
		fmt.Println(parent)
	}

	return dist
}

func main() {
	graph := make([][]Edge, 5)

	graph[0] = []Edge{{to: 1, weight: 10}, {to: 2, weight: 3}}
	graph[1] = []Edge{{to: 2, weight: 1}, {to: 3, weight: 2}}
	graph[2] = []Edge{{to: 1, weight: 4}, {to: 3, weight: 8}, {to: 4, weight: 2}}
	graph[3] = []Edge{{to: 4, weight: 7}}
	graph[4] = []Edge{{to: 3, weight: 9}}

	startNode := 0
	distances := Dijkstra(graph, startNode)

	fmt.Printf("Кратчайшие расстояния от узла %d:\n", startNode)
	for node, d := range distances {
		fmt.Printf("До узла %d: %d\n", node, d)
	}
}