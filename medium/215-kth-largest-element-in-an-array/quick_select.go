package KthLargestElement

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())

	var quickSelect func(sub []int, p int) int
	quickSelect = func(sub []int, p int) int {
		pivot := rand.Intn(len(sub))

		left := []int{}
		right := []int{}
		mid := []int{}

		for _, num := range sub {
			if num > sub[pivot] {
				left = append(left, num)
			}
			if num < sub[pivot] {
				right = append(right, num)
			}
			if num == sub[pivot] {
				mid = append(mid, num)
			}
		}

		if p <= len(left) {
			return quickSelect(left, p)
		}
		if p > len(left)+len(mid) {
			return quickSelect(right, p-(len(left)+len(mid)))
		}

		return sub[pivot]
	}

	return quickSelect(nums, k)
}
