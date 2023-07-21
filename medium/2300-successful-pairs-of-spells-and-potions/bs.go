package SuccessfulPairs

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	ns := len(spells)
	np := len(potions)
	sort.Ints(potions)

	var bst func(goal float64) int
	bst = func(goal float64) int {
		l := 0
		e := np
		m := np / 2
		for e-l > 1 {
			if goal <= float64(potions[m]) {
				e = m
				m = (e + l) / 2
			} else {
				l = m
				m = (e + l) / 2
			}
		}

		if float64(potions[l]) >= goal {
			return l
		}
		if l+1 < np {
			return l + 1
		}
		return -1
	}

	res := make([]int, ns)
	for j := 0; j < ns; j++ {
		goal := float64(success) / float64(spells[j])
		r := bst(goal)
		if r < 0 {
			res[j] = 0
		} else {
			res[j] = np - r
		}
	}
	return res
}
