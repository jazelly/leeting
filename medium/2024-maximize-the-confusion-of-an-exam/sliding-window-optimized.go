package main

// TODO: revisit to grasp the intention
func maxConsecutiveAnswers(answerKey string, k int) int {
	rF, lF, kF := 0, 0, k
	rT, lT, kT := 0, 0, k
	for rF < len(answerKey) {
		if answerKey[rF] == 'F' {
			kF--
		}

		if answerKey[rT] == 'T' {
			kT--
		}

		if kF < 0 {
			if answerKey[lF] == 'F' {
				kF++
			}
			lF++
		}

		if kT < 0 {
			if answerKey[lT] == 'T' {
				kT++
			}
			lT++
		}

		rF++
		rT++
	}
	return max(rT-lT, rF-lF)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxConsecutiveAnswers("TTFTTTFT", 1)
}
