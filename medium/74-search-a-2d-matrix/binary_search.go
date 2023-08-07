package searchMatrix

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	l := 0
	r := m
	for r-l > 1 {
		mid := (l + r) / 2
		if target >= matrix[mid][0] {
			l = mid
		} else {
			r = mid
		}
	}

	ll := 0
	n := len(matrix[l])
	rr := n
	for rr-ll > 1 {
		mid := (ll + rr) / 2
		if target >= matrix[l][mid] {
			ll = mid
		} else {
			rr = mid
		}
	}

	return matrix[l][ll] == target
}
