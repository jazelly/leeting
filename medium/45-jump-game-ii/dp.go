package jump2

func jump(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	for i := 1; i < n; i++ {
		dp[i] = 20000
	}

	for i := 0; i < n; i++ {
		for j := i; j <= i+nums[i] && j < n; j++ {
			if dp[i]+1 < dp[j] {
				dp[j] = dp[i] + 1
			}
		}
	}

	return dp[n-1]
}
