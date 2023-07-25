package PeakIndexInMountainArray

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)

	l := 0
	r := n

	for r-l > 1 {
		m := (r + l) / 2

		if arr[m] < arr[m+1] {
			l = m
		} else {
			r = m
		}
	}

	if arr[r] > arr[l] {
		return r
	} else {
		return l
	}
}
