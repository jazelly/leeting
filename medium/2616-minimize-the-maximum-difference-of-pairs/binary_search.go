package MinimizeMax

import "sort"

func minimizeMax(nums []int, p int) int {
	n := len(nums)
	sort.Ints(nums)
	if p == 0 {
		return 0
	}

	countPairs := func(t int) int {
		k := 0
		for i := 0; i < n-1; {
			if nums[i+1]-nums[i] <= t {
				k++
				i++
			}
			i++
		}
		return k
	}

	l := 0
	r := nums[n-1] - nums[0]
	for r > l {
		m := (r + l) / 2
		if countPairs(m) >= p {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
