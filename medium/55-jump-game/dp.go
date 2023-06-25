package jump

func canJump(nums []int) bool {
	// the best I can reach
	farthest := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if nums[i] == 0 && i < n-1 && i == farthest {
			return false
		}

	}

	return true
}
