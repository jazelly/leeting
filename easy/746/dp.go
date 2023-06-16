package main

// Input: cost = [10,15,20]
// Output: 15
// Explanation: You will start at index 1.
// - Pay 15 and climb two steps to reach the top.
// The total cost is 15.

func minCostClimbingStairs(cost []int) int {
	// [10, 15, 20, 10, 30, 20, 10]
	// [0,  0, INF,INF,INF,INF,INF]
	// [0,  0,  10, 15, 25]
	n := len(cost)
	if n == 1 {
		return 0
	}
	if n == 2 {
		if cost[0] < cost[1] {
			return cost[0]
		} else {
			return cost[1]
		}
	}
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < n; i++ {
		two := dp[i-2] + cost[i-2]
		one := dp[i-1] + cost[i-1]
		if two < one {
			dp[i] = two
		} else {
			dp[i] = one
		}
	}

	one := dp[n-1] + cost[n-1]
	two := dp[n-2] + cost[n-2]

	if one < two {
		return one
	} else {
		return two
	}
}

func main() {
	inp := []int{10, 15, 20, 10, 30, 20, 10}

	minCostClimbingStairs(inp)
}
