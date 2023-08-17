package UpdateMatrix

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	visited := make([][]int, m)
	q := [][]int{}
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j, 0})
			}
		}
	}

	dir := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for len(q) > 0 {
		head := q[0]
		q = q[1:]

		y := head[0]
		x := head[1]
		dis := head[2]
		if visited[y][x] == 1 {
			continue
		}

		for _, d := range dir {
			newY := y + d[0]
			newX := x + d[1]
			if newY < 0 || newX < 0 || newY >= m || newX >= n {
				continue
			}
			if mat[newY][newX] == 0 {
				continue
			}
			if visited[newY][newX] == 1 {
				continue
			}

			q = append(q, []int{newY, newX, dis + 1})
		}

		mat[y][x] = dis
		visited[y][x] = 1
	}

	return mat
}
