package soupServings

func soupServings(n int) float64 {
	if n > 4800 {
			return 1
	}
	np := n/25
	if n%25 > 0 {
			np++
	}
	np++

	// a * b
	// ath serve of a and bth serve of b
	v := make([][]float64, np)
	for i:=0; i<np;i++ {
			v[i] = make([]float64, np)
			for j :=0; j < np;j++ {
					v[i][j] = -1
			}
	}

	for i := 0; i<np; i++ {
			v[np-1][i] = 1
	}
	for i := 0; i<np; i++ {
			v[i][np-1] = 0
	}
	v[np-1][np-1] = 0.5


	var dfs func(a int, b int, ai int, bi int) float64
	dfs = func(a int, b int, ai int, bi int) float64 {
			if a <= 0 && b <= 0 {
					if ai < np && bi < np { v[ai][bi] = 0.5 }
					return 0.5
			}
			if a <= 0 && b > 0 {
					if ai < np && bi < np { v[ai][bi] = 1 }
					return 1
			}
			if a > 0 && b <= 0 {
					if ai < np && bi < np { v[ai][bi] = 0 }
					return 0
			} 

			var r float64
			if ai + 4 < np && v[ai+4][bi] >= 0 {
					r += v[ai+4][bi]
			} else {
					r += dfs(a-100, b, ai+4, bi)
			}
			if ai + 3 < np && bi + 1 < np && v[ai+3][bi+1] >= 0 {
					r += v[ai+3][bi+1]
			} else {
					r += dfs(a-75, b-25, ai+3, bi+1)
			}
			if ai + 2 < np && bi + 2 < np && v[ai+2][bi+2] >= 0 {
					r += v[ai+2][bi+2]
			} else {
					r += dfs(a-50, b-50, ai+2, bi+2)
			}
			if ai + 1 < np && bi + 3 < np && v[ai+1][bi+3] >= 0 {
					r += v[ai+1][bi+3]
			} else {
					r += dfs(a-25, b-75, ai+1, bi+3)
			}

			r *= 0.25
			v[ai][bi] = r
			return r
	}

	return dfs(n, n, 0, 0)
}
