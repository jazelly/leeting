package FindTheLongestEqualSubarray

func longestEqualSubarray(nums []int, k int) int {
	n := len(nums)
	l := 0
	r := 0

	wmap := make(map[int]int)
	wmax := make(map[int]int)
	tMax := 0
	max := 0
	for r < n {
		wmap[nums[r]]++
		if wmap[nums[r]] >= tMax {
			tMax = wmap[nums[r]]
			wmax[nums[r]] = tMax
		}

		if tMax > max {
			max = tMax
		}

		cover := r - l + 1
		if cover-tMax-k > 0 {
			if wmap[nums[l]] == tMax {
				delete(wmax, nums[l])
			}
			wmap[nums[l]]--
			l++
		}
		r++
	}
	return max
}
