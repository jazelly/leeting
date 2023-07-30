package minimumString

func minimumString(a string, b string, c string) string {
	res := "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
	
	keep := func(abcabc string) {
			if len(abcabc) < len(res) {
					res = abcabc
					return
			}
			if len(abcabc) == len(res) {
					for i:=0; i< len(res); i++ {
							if abcabc[i] < res[i] {
									res = abcabc
									return
							} else if abcabc[i] == res[i] {
									continue
							} else {
									break
							}
					}
			}
	}

	keep(findOverlap(c, findOverlap(a, b)))
	keep(findOverlap(c, findOverlap(b, a)))
	keep(findOverlap(a, findOverlap(b, c)))
	keep(findOverlap(a, findOverlap(c, b)))
	keep(findOverlap(b, findOverlap(a, c)))
	keep(findOverlap(b, findOverlap(c, a)))
	keep(findOverlap(findOverlap(a, b), c))
	keep(findOverlap(findOverlap(b, a), c))
	keep(findOverlap(findOverlap(a, c), b))
	keep(findOverlap(findOverlap(c, a), b))
	keep(findOverlap(findOverlap(b, c), a))
	keep(findOverlap(findOverlap(c, b), a))

	return res
}


func findOverlap(a string, b string) string {
	if strings.Contains(a, b) {
			return a
	}
	
	// a + b
	var ab string
	i := 0
	j := len(a) - 1
	
	resi := 0
	for i < len(b) && j >= 0 {
			if a[j:] == b[:i+1] {
					resi = i+1
			}
			i++
			j-- 
	}
	
	ab = a + b[resi:]    
	
	return ab
}