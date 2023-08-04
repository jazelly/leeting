package wordBreak

func wordBreakBFS(s string, wordDict []string) bool {
	n := len(s)

	dd := make(map[byte][]string)
	for _, w := range wordDict {
		dd[w[0]] = append(dd[w[0]], w)
	}

	// bfs
	q := []int{0}
	for len(q) > 0 {
		newQ := []int{}
		visited := make(map[int]bool)

		for _, p := range q {
			for _, w := range dd[s[p]] {
				le := len(w)
				r := p + le
				if visited[r] {
					continue
				}
				if r > n {
					continue
				}
				if s[p:r] == w {
					if r == n {
						return true
					}
					visited[r] = true
					newQ = append(newQ, r)
				}
			}
		}

		q = newQ
	}

	return false
}
