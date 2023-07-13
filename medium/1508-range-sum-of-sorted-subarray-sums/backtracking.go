package rangeSum

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	res := []int{}

	for i := 0; i < n; i++ {
		buffer := 0
		for j := 0; i+j < n; j++ {
			buffer += nums[i+j]
			if buffer > 0 {
				res = append(res, buffer)
			}
		}
	}

	sort.Ints(res)

	sum := 0
	for i := left - 1; i < right; i++ {
		sum = sum + res[i]
	}

	var number int = 1000000007
	return sum % number
}
