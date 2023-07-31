package countCompleteSubarrays

func countCompleteSubarrays(nums []int) int {
	dwmap := make(map[int]bool)
	for _, v := range nums {
		dwmap[v] = true
	}

	dw := len(dwmap)
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		wmap := make(map[int]int)
		for j := i; j < n; j++ {
			wmap[nums[j]]++
			if len(wmap) == dw {
				res++
			}
		}
	}
	return res
}
