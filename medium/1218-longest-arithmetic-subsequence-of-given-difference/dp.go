package longestSubsequence

func longestSubsequence(arr []int, difference int) int {
	n := len(arr)
	dp := make(map[int]int, n)

	for i := 0; i < n; i++ {
		if dp[arr[i]-difference] > 0 {
			dp[arr[i]] = dp[arr[i]-difference] + 1
		} else {
			dp[arr[i]] = 1
		}
	}

	m := 0
	for _, v := range dp {
		if v > m {
			m = v
		}
	}

	return m
}
