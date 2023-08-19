package Dota2Senate

func predictPartyVictory2(senate string) string {
	q := []byte{}
	n := len(senate)
	for i := 0; i < len(senate); i++ {
		q = append(q, senate[i])
	}

	i := 0
	r := 0
	d := 0
	for i < 2*n {
		if q[0] == byte('R') {
			if d == 0 {
				r++
				q = append(q, q[0])
			} else {
				d--
			}
		} else {
			if r == 0 {
				d++
				q = append(q, q[0])
			} else {
				r--

			}
		}
		q = q[1:]
		i++
	}
	if q[0] == byte('D') {
		return "Dire"
	}
	return "Radiant"
}
