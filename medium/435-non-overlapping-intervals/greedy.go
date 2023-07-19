package eraseOverlapIntervals

import "sort"

// classic meeting schedule problem
// classic greedy
// we aim to maximize the intervals we can have
// so for every meeting finished, we aim to leave max space for the rest
func eraseOverlapIntervals(intervals [][]int) int {
	lessFunc := func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	}

	sort.Slice(intervals, lessFunc)

	n := len(intervals)
	prev := 0
	c := 1

	for i := 1; i < n; i++ {
		if intervals[i][0] >= intervals[prev][1] {
			prev = i
			c++
		}
	}
	return n - c
}
