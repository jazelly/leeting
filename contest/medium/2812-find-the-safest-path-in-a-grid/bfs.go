package FindTheSafestPath

// https://leetcode.com/contest/weekly-contest-357/problems/find-the-safest-path-in-a-grid/
func maximumSafenessFactor(grid [][]int) int {
	ts := [][]int{} // thieves
	n := len(grid)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ts = append(ts, []int{i, j})
			}
		}
	}

	// update to distance-to-thieves map
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			min := 100000000
			for _, t := range ts {
				if abs(t[0]-i)+abs(t[1]-j) < min {
					min = abs(t[0]-i) + abs(t[1]-j)
				}
			}
			grid[i][j] = min
		}

	}

	// just a normal bfs to check if the path can meet the requirement r
	// i.e. give a safenessFactor r, see if any path can meet it
	var bfsCheck func(r int) bool
	bfsCheck = func(r int) bool {
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}

		q := [][]int{}
		q = append(q, []int{0, 0})
		for len(q) > 0 {
			front := q[0]
			x := front[0]
			y := front[1]

			if x < 0 || x >= n || y < 0 || y >= n || visited[x][y] || grid[x][y] < r {
				q = q[1:]
				continue
			}

			q = append(q, []int{x - 1, y})
			q = append(q, []int{x, y - 1})
			q = append(q, []int{x + 1, y})
			q = append(q, []int{x, y + 1})
			visited[x][y] = true
			q = q[1:]
		}

		return visited[n-1][n-1]
	}

	l := 0
	r := 401
	for r-l > 1 {
		m := (l + r) / 2
		if bfsCheck(m) {
			l = m
		} else {
			r = m
		}
	}
	return l
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
