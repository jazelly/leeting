package letterCombinations

func letterCombinations(digits string) []string {
	phoneMap := map[int][]rune{
		2: []rune{'a', 'b', 'c'},
		3: []rune{'d', 'e', 'f'},
		4: []rune{'g', 'h', 'i'},
		5: []rune{'j', 'k', 'l'},
		6: []rune{'m', 'n', 'o'},
		7: []rune{'p', 'q', 'r', 's'},
		8: []rune{'t', 'u', 'v'},
		9: []rune{'w', 'x', 'y', 'z'},
	}

	n := len(digits)
	res := []string{}
	var dfs func(d int, buffer string)
	dfs = func(d int, buffer string) {
		if d == n {
			if len(buffer) > 0 {
				res = append(res, buffer)
			}
			return
		}

		v := int(digits[d] - '0')
		for _, o := range phoneMap[v] {
			dfs(d+1, buffer+string(o))
		}
	}
	dfs(0, "")

	return res
}
