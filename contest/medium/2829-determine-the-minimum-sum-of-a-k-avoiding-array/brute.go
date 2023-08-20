package MinimumSum

func minimumSum(n int, k int) int {
	res := []int{}
	v := 1
	twosum := make(map[int]bool)
	for len(res) < n {
		if !twosum[v] {
			twosum[k-v] = true
			res = append(res, v)
		}
		v++
	}
	sum := 0
	for _, r := range res {
		sum += r
	}

	return sum
}
