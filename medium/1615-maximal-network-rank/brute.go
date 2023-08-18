package MaximalNetworkRank

func maximalNetworkRank(n int, roads [][]int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	for _, v := range roads {
		graph[v[0]][v[1]] = 1
		graph[v[1]][v[0]] = 1
	}

	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != i {
				cij := 0
				for _, v := range graph[i] {
					cij += v
				}

				for _, v := range graph[j] {
					cij += v
				}

				if graph[i][j] == 1 {
					cij--
				}

				if cij > max {
					max = cij
				}
			}

		}
	}
	return max
}
