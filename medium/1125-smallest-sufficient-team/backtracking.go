package smallestsufficientteam

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	rmap := make(map[string]int)
	for i, v := range req_skills {
		rmap[v] = i
	}

	var target uint = (1 << len(req_skills)) - 1
	min := 100
	var realAns []int

	// dfs
	var dfs func(i int, incoming uint, target uint, ans []int, rmap map[string]int, people [][]string)
	dfs = func(i int, incoming uint, target uint, ans []int, rmap map[string]int, people [][]string) {
		if len(ans) >= min {
			return
		}

		if incoming == target {
			if len(ans) < min {
				realAns = make([]int, len(ans))
				min = len(ans)
				copy(realAns, ans)
			}
			return
		}

		if i == len(people) {
			return
		}

		var gain uint = 0
		for _, ps := range people[i] {
			gain |= (1 << rmap[ps])
		}

		pick := gain | incoming

		if pick > incoming {
			ans = append(ans, i)
			dfs(i+1, pick, target, ans, rmap, people)
			ans = ans[:len(ans)-1]
		}

		dfs(i+1, incoming, target, ans, rmap, people)

		return
	}

	dfs(0, 0, target, []int{}, rmap, people)

	return realAns
}
