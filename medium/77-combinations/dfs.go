package combinations

func combine(n int, k int) [][]int {
	var dfs func(j int, l int, buffer []int)
	res := [][]int{}
	dfs = func(j int, l int, buffer []int) {
		if l >= k {
			c := make([]int, len(buffer))
			copy(c, buffer)
			res = append(res, c)
			return
		}

		for i := j; i <= n-k+l+1; i++ {
			added := append(buffer, i)
			dfs(i+1, l+1, added)
		}
	}

	dfs(1, 0, []int{})

	return res
}
