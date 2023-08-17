package MaxSlidingWindow

func maxSlidingWindowBrute(nums []int, k int) []int {
	maxw := -100000
	maxi := -1

	find := func(l, r int) {
		maxw = -100000
		maxi = -1
		for i := l; i <= r; i++ {
			if nums[i] > maxw {
				maxw = nums[i]
				maxi = i
			}
		}
	}
	find(0, k-1)

	n := len(nums)
	res := make([]int, n-k+1)
	res[0] = maxw
	c := 1
	for i := k; i < n; i++ {
		popi := i - k
		if maxi == popi {
			find(i-k+1, i)
		} else if nums[i] > maxw {
			maxw = nums[i]
			maxi = i
		}
		res[c] = maxw
		c++
	}
	return res
}
