package maxConsecutiveAnswers

func maxConsecutiveAnswers(answerKey string, k int) int {
	if len(answerKey) < 2 {
		return len(answerKey)
	}

	n := len(answerKey)
	l, r, res, fmap := 0, 0, 0, make(map[byte]int, 2)
	fmap[byte(84)] = 0
	fmap[byte(70)] = 0

	for ; r < n; r++ {
		fmap[answerKey[r]]++

		flipCount := getSmallestValue(fmap)
		for flipCount > k {
			fmap[answerKey[l]]--
			l++
			flipCount = getSmallestValue(fmap)
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}

func getSmallestValue(fmap map[byte]int) int {
	min := 100000
	for _, v := range fmap {
		if v < min {
			min = v
		}
	}

	return min
}
