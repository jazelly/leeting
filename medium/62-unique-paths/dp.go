// https://leetcode.com/problems/unique-paths/

package uniquePaths

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		row := make([]int, n)
		dp[i] = row
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1]
}
