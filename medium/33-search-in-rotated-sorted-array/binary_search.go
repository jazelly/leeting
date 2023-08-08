package SearchInRotatedSortedArray

func search(nums []int, target int) int {
	n := len(nums)
	pivot := findPivot(nums, n)
	var low int
	var high int

	if pivot < 0 {
		low = 0
		high = n - 1
	} else if target < nums[0] {
		low = pivot
		high = n - 1
	} else {
		low = 0
		high = pivot - 1
	}

	for low < high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if nums[low] == target {
		return low
	}
	return -1
}

func findPivot(nums []int, n int) int {
	low := 0
	high := n - 1
	if nums[low] < nums[high] {
		return -1
	}

	for low < high {
		mid := (low + high) / 2

		// left part is normal
		if nums[mid] >= nums[0] {
			low = mid + 1
		} else {
			// pivot is at left part
			high = mid
		}
	}

	return low
}
