package main

func distributeCookies(cookies []int, k int) int {
	ans := 8000000
	children := make([]int, k)

	var dfs func(int)
	dfs = func(curCookie int) {
		if curCookie == len(cookies) {
			mx := 0
			for i := 0; i < k; i++ {
				mx = max(mx, children[i])
			}
			ans = min(mx, ans)
			return
		}

		for i := 0; i < k; i++ {
			if children[i]+cookies[curCookie] > ans {
				continue
			}
			children[i] += cookies[curCookie]
			dfs(curCookie + 1)
			children[i] -= cookies[curCookie]
		}
	}

	dfs(0)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cookies := []int{
		8, 10, 15, 8, 20,
	}
	distributeCookies(cookies, 2)
}
