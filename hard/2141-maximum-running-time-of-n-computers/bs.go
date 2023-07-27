package maxRunTime

func maxRunTime(n int, batteries []int) int64 {
	l := 1
	s := 0
	for _, v := range batteries {
		s += v
	}
	r := s/n + 1
	var target int
	for r-l > 1 {
		target = (r + l) / 2
		allPowerForTarget := 0
		for i := 0; i < len(batteries); i++ {
			if batteries[i] > target {
				allPowerForTarget += target
			} else {
				allPowerForTarget += batteries[i]
			}
		}
		if allPowerForTarget < n*target {
			r = target
		} else {
			l = target
		}
	}

	return int64(l)
}
