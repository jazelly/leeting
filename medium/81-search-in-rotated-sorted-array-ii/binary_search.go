package SearchInRotatedSortedArrayII

func search(nums []int, target int) bool {
	n := len(nums)
	l := 0
	r := n
	// binary search the original ending point
	for r-l > 1 {
		m := (r + l) / 2
		if nums[m] == nums[l] {
			l++
			continue
		}

		if nums[m] > nums[l] {
			l = m
		} else {
			r = m
		}
	}

	start := l + 1
	left := nums[:start]
	l = 0
	r = len(left)
	for r-l > 1 {
		m := (r + l) / 2
		if left[m] == target {
			return true
		}
		if left[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if l < len(left) && left[l] == target {
		return true
	}

	right := nums[start:]
	l = 0
	r = len(right)
	for r-l > 1 {
		m := (r + l) / 2
		if right[m] == target {
			return true
		}
		if right[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if l < len(right) && right[l] == target {
		return true
	}
	return false
}
