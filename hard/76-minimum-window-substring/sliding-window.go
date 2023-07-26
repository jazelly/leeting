package minWindow

func minWindow(s string, t string) string {
	tmap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++
	}
	wmap := make(map[byte]int)
	n := len(s)
	res := n + n
	ress := []int{0, 0}

	l := 0
	r := l

	valid := 0

	for r < n {
		wmap[s[r]]++
		if tmap[s[r]] > 0 && wmap[s[r]] == tmap[s[r]] {
			valid++
		}
		r++
		for valid == len(tmap) {
			if r-l < res {
				res = r - l
				ress[0] = l
				ress[1] = r
			}

			// move l
			tbr := s[l]
			l++
			if tmap[tbr] > 0 && wmap[tbr] == tmap[tbr] {
				valid--
			}
			wmap[tbr]--
		}
	}

	if res == n+n {
		return ""
	}
	return s[ress[0]:ress[1]]
}
