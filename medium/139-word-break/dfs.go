package wordBreak

import "sort"

// This approach will TLE
func wordBreak(s string, wordDict []string) bool {
	sort.Slice(wordDict, func(i, j int) bool {
		return len(wordDict[i]) > len(wordDict[j])
	})

	dd := make(map[byte][]string)
	for _, w := range wordDict {
		dd[w[0]] = append(dd[w[0]], w)
	}

	n := len(s)
	var dfs func(p int) bool
	dfs = func(p int) bool {
		if p == n {
			return true
		}
		if len(dd[s[p]]) == 0 {
			return false
		}

		res := false
		for _, w := range dd[s[p]] {
			le := len(w)
			r := p + le
			if r > n {
				continue
			}
			if s[p:r] == w {
				if dfs(r) {
					res = true
					break
				}
			}
		}

		return res
	}

	return dfs(0)
}
