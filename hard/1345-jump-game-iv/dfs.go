package JumpGameIV

// This approach will be TLE for large input
func minJumpsDFS(arr []int) int {
	position := make(map[int][]int)

	n := len(arr)
	for i, v := range arr {
		position[v] = append(position[v], i)
	}

	dp := make([]int, n)

	var dfs func(d int, p int)
	dfs = func(d int, p int) {
		if p == n-1 {
			if d < dp[p] && dp[p] > 0 {
				dp[p] = d
			} else if dp[p] == 0 {
				dp[p] = d
			}
			return
		}

		if d >= dp[p] && dp[p] > 0 {
			return
		}

		options := make(map[int]bool)
		if p-1 >= 1 {
			options[p-1] = true
		}
		if p+1 < n {
			options[p+1] = true
		}
		for _, v := range position[arr[p]] {
			options[v] = true
		}

		dp[p] = d

		for o, _ := range options {
			dfs(d+1, o)
		}
	}
	dfs(0, 0)

	return dp[n-1]
}
