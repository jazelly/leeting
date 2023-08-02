package permutations

func permute2(nums []int) [][]int {
	res := [][]int{}
	var dfs func(options []int, buffer []int)
	dfs = func(options []int, buffer []int) {
		if len(options) == 0 {
			c := make([]int, len(buffer))
			copy(c, buffer)
			res = append(res, c)
		}

		for i, v := range options {
			added := append(buffer, v)
			cut := []int{}
			cut = append(cut, options[:i]...)
			if i < len(options)-1 {
				cut = append(cut, options[i+1:]...)
			}
			dfs(cut, added)
		}
	}

	dfs(nums, []int{})
	return res
}
