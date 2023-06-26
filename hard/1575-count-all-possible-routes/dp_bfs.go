// https://leetcode.com/problems/count-all-possible-routes/

package countAllPossibleRoutes

func countRoutes(locations []int, start int, finish int, fuel int) int {
	// the discovery of routes is like discovering a tree starting from root start, with full fuel
	// the results is actually the sum of all possible amount of left fuel at destination
	// so we just need to decrement the fuel while discovering the routes
	// and cache them in our location-fuel table

	// this is BFS + caching, aka memorisation, aka DP
	n := len(locations)
	mod := int(1e9) + 7
	dp := make([][]int, n)

	// usages(0) = [0, 1, 4, 6, 2]
	// usages(1) = [1, 0, 3, 5, 1]
	// usages(2) = [4, 3, 0, 2, 2]
	// usages(3) = [6, 5, 2, 0, 4]
	// usages(4) = [2, 1, 2, 4, 0]
	usages := make([][]int, n)
	for i := 0; i < n; i++ {
		usage := make([]int, n)
		for j := 0; j < n; j++ {
			usage[j] = diff(locations[i], locations[j])
		}
		usages[i] = usage
	}

	// e.g. Example 1
	// starting from full fuel 5, l = start
	// we can reach out to
	// l  f0, 1, 2, 3, 4, 5
	// 0  [0, 0, 0, 0, 0, 0]
	// 1  [0, 0, 0, 0, 0, 1]
	// 2  [0, 0, 0, 0, 0, 0]
	// 3  [0, 0, 0, 0, 0, 0]  sum() is result
	// 4  [0, 0, 0, 0, 0, 0]
	for i := 0; i < n; i++ {
		fuelLeft := make([]int, fuel+1)
		dp[i] = fuelLeft
	}
	dp[start][fuel]++

	for f := fuel; f >= 0; f-- {
		for l := 0; l < n; l++ {
			lf := dp[l][f]
			if lf > 0 {
				// for each usage in the usages map
				for i := 0; i < n; i++ {
					if usages[l][i] > 0 && usages[l][i] <= f && l != i {
						dp[i][f-usages[l][i]] += lf
						if dp[i][f-usages[l][i]] > mod {
							dp[i][f-usages[l][i]] = dp[i][f-usages[l][i]] % mod
						}
					}
				}
			}
		}
	}

	sum := 0
	for i := 0; i <= fuel; i++ {
		sum += dp[finish][i]
	}
	return sum % mod
}

func diff(a int, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
