// https://leetcode.com/problems/unique-paths-ii/

package uniquePaths2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 0 {
		obstacleGrid[0][0] = 1
	} else {
		obstacleGrid[0][0] = 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				if j > 0 {
					obstacleGrid[i][j] += obstacleGrid[i][j-1]
				}
				if i > 0 {
					obstacleGrid[i][j] += obstacleGrid[i-1][j]
				}
			}
		}
	}

	return obstacleGrid[m-1][n-1]
}
