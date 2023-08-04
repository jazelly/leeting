package JumpGameIV

func minJumps(arr []int) int {
	n := len(arr)
	position := make(map[int][]int)

	for i, v := range arr {
		position[v] = append(position[v], i)
	}

	d := 0
	q := []int{0}
	visited := make(map[int]bool, n)
	for len(q) > 0 {
		// deduplicate next layer
		options := make(map[int]bool)
		for _, p := range q {
			if p == n-1 {
				return d
			}
			if p-1 >= 1 {
				options[p-1] = true
			}
			if p+1 < n {
				options[p+1] = true
			}
			for _, v := range position[arr[p]] {
				options[v] = true
			}
			visited[p] = true

			// this is very important
			// once we have pushed same value positions
			// we won't use it again, as this list could be very large
			delete(position, arr[p])
		}

		newQ := []int{}
		for o, ok := range options {
			if o == n-1 {
				return d + 1
			}
			if !visited[o] && ok {
				newQ = append(newQ, o)
			}
		}
		q = newQ
		d++
	}

	return d
}
