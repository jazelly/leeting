package ValidPartition

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)

	checkValid := func(j, k int) {
		if nums[j] == nums[j+1] {
			dp[j+1] = 1
		}
		if k-j == 1 {
			return
		}

		if nums[j] == nums[j+1] && nums[j+1] == nums[j+2] {
			dp[j+2] = 1
		}

		if nums[j]+1 == nums[j+1] && nums[j+1]+1 == nums[j+2] {
			dp[j+2] = 1
		}
	}

	checkValid(0, 1)
	if n == 2 {
		return dp[n-1] == 1
	}

	checkValid(0, 2)
	for i := 3; i < n; i++ {
		if dp[i-2] == 1 {
			checkValid(i-1, i)
		}

		if dp[i-3] == 1 {
			checkValid(i-2, i)
		}
	}
	return dp[n-1] == 1
}
