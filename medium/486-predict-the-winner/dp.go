package PredictTheWinner

// Please visit https://leetcode.com/problems/predict-the-winner/editorial/

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i:=0; i < n; i++ {
			dp[i] = make([]int, n)
			dp[i][i] = nums[i]
	}

	for i := 1; i < n; i++ {
			for j := 0; j < n-i; j++ {
					l := nums[j] - dp[j+1][j+i]
					r := nums[j+i] - dp[j][j+i-1]
					dp[j][j+i] = max(l, r)
			}
	}

	return dp[0][n-1] >= 0
}

func max(a int, b int) int {
	if a > b {
			return a
	}
	return b
}
