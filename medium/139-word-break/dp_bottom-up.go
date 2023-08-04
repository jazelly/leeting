package wordBreak

func wordBreakDP(s string, wordDict []string) bool {
	n := len(s)

	dd := make(map[string]bool)
	for _, w := range wordDict {
		dd[w] = true
	}

	dp := make([]int, n)

	p := []int{-1}
	for i := 0; i < n; i++ {
		for _, j := range p {
			if dd[s[j+1:i+1]] {
				dp[i] = 1
				p = append(p, i)
				break
			}
		}
	}
	return dp[n-1] == 1
}
