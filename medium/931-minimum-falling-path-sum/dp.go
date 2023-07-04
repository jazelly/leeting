package minFallingPathSum

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			min := 1000000
			for k := j - 1; k <= j+1; k++ {
				if k >= 0 && k <= n-1 {
					if matrix[i-1][k]+matrix[i][j] < min {
						min = matrix[i-1][k] + matrix[i][j]
					}
				}
			}
			matrix[i][j] = min
		}
	}

	min := 10000000
	for j := 0; j < n; j++ {
		if matrix[n-1][j] < min {
			min = matrix[n-1][j]
		}
	}
	return min
}
