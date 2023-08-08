package FindTheSafestPath

import (
	"container/heap"
	"fmt"
)

func maximumSafenessFactorDijkstra(grid [][]int) int {
	ts := [][]int{} // thieves
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}
	safe := make([][]int, n)
	for i := 0; i < n; i++ {
		safe[i] = make([]int, n)
		for j := 0; j < n; j++ {
			safe[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ts = append(ts, []int{i, j})
				safe[i][j] = 0
			}
		}
	}

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, -1},
		{0, 1},
	}

	// update to distance-to-thieves map
	// multi-source bfs
	for len(ts) > 0 {
		newTs := [][]int{}
		for _, t := range ts {
			for _, d := range directions {
				ntx := t[0] + d[0]
				nty := t[1] + d[1]

				if ntx < 0 || ntx >= n || nty < 0 || nty >= n || safe[ntx][nty] >= 0 {
					continue
				}
				safe[ntx][nty] = safe[t[0]][t[1]] + 1
				newTs = append(newTs, []int{ntx, nty})
			}
		}
		ts = newTs
	}
	fmt.Println(safe)
	hp := &MaxHeap{}
	heap.Init(hp)
	heap.Push(hp, []int{0, 0, safe[0][0]})

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	for hp.Len() > 0 {
		node := hp.Pop().([]int)
		if node[0] == n-1 && node[1] == n-1 {
			return node[2]
		}
		for _, d := range directions {
			ntx := node[0] + d[0]
			nty := node[1] + d[1]
			if ntx < 0 || ntx >= n || nty < 0 || nty >= n || visited[ntx][nty] {
				continue
			}
			hp.Push([]int{ntx, nty, min(node[2], safe[ntx][nty])})
			visited[ntx][nty] = true
		}
	}

	return 7474741
}

type MaxHeap [][]int

// Implement the heap.Interface methods for MaxHeap
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][2] > h[j][2] } // Max-heap property
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

// Pop removes and returns the maximum element from the heap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}
