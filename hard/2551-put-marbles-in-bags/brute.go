package putMarbles

import "sort"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)

	if k == 1 {
		return int64(0)
	}

	// [4,3,5,1,2,4,5,6,6,7,1]
	// we are just looing the max sum of 2 consecutive nums of the array
	// and only k-1 of them
	// i.e.
	// [7 8 6 3 6 9 11 12 13 8]
	max := 0
	min := 0

	think := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		think[i] = weights[i] + weights[i+1]
	}

	sort.Ints(think)

	for i := 0; i < k-1; i++ {
		min += think[i]
	}
	for i := n - 2; i > n-1-k; i-- {
		max += think[i]
	}

	return int64(max - min)
}
