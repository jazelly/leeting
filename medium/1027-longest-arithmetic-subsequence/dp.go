package longestArithSeqLength

func longestArithSeqLength(nums []int) int {
	n := len(nums)

	dp := make([]map[int]int, n)
	dp[0] = make(map[int]int, 0)

	res := 0
	for i := 1; i < n; i++ {
		diffMap := make(map[int]int, i)
		dp[i] = diffMap
		for j := 0; j < i; j++ {
			aDiff := nums[i] - nums[j]
			if val, ok := dp[j][aDiff]; ok {
				dp[i][aDiff] = val + 1
			} else {
				dp[i][aDiff] = 1
			}

			if dp[i][aDiff] > res {
				res = dp[i][aDiff]
			}
		}
	}

	res++
	return res
}
