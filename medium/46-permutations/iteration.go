package permutations

func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	r := n

	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	cycles := make([]int, r)
	for i := 0; i < n; i++ {
		cycles[i] = r - i
	}

	res = append(res, yield(nums, r, indices))
	for true {
		broke := false
		for i := n - 1; i >= 0; i-- {
			cycles[i] -= 1
			if cycles[i] == 0 {
				head := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = head
				cycles[i] = n - i
			} else {
				ci := cycles[i]
				indices[i], indices[n-ci] = indices[n-ci], indices[i]
				res = append(res, yield(nums, r, indices))
				broke = true
				break
			}
		}
		if !broke {
			return res
		}
	}
	return res
}

func yield(nums []int, r int, indices []int) []int {
	tmp := make([]int, r)
	for i := 0; i < r; i++ {
		tmp[i] = nums[indices[i]]
	}
	return tmp
}
