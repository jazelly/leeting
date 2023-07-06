package minSubArrayLen

import "math"

func minSubArrayLen(target int, nums []int) int {
	cumulative := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		cumulative[i+1] = cumulative[i] + nums[i]
	}

	i := 0
	j := 1

	minlen := math.MaxInt

	for j < len(cumulative) {
		diff := cumulative[j] - cumulative[i]
		if diff < target {
			j++
		} else {
			if j-i < minlen {
				minlen = j - i
			}
			i++
		}
	}

	if minlen == math.MaxInt {
		return 0
	}

	return minlen
}
