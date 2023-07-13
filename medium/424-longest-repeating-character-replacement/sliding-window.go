package main

import "fmt"

func characterReplacement(s string, k int) int {
	n := len(s)
	l := 0

	windowMap := make(map[byte]int)
	theMax := 0
	maxw := 0
	for i := 0; i < n; i++ {
		windowMap[s[i]]++
		if windowMap[s[i]] >= maxw {
			maxw = windowMap[s[i]]
			if maxw > theMax {
				theMax = maxw
			}
		}

		if i-l+1 > maxw+k {
			windowMap[s[l]]--
			l++
		}

	}
	if theMax+k > n {
		return n
	}

	return theMax + k
}

func main() {

	s := "BRJRRKNRBFOOKDEEGODTGMHNABMTHFNPTFRHRSEKKTFEQIKSIAJJMSDSLNSCNRNJFNFSIQDNMHDRIJIACLCJKATTFHDASGLRQSFN"
	fmt.Println(string(s[60]))
	characterReplacement(s, 10)

}
