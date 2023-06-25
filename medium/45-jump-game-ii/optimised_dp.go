// https://leetcode.com/problems/jump-game-ii/description/

// O(n) time and O(1) space
// all we need to remember is what is the farthest index we can reach
// and who brought us there
// but, since the question doesn't really care about the path
// we don't need to remember who exactly is the predecessor
// we just need to know what is the min cost brought us to the farthest

// To know the min cost to the farthest
// we just need to know the min cost to the current window
// in other word, unless we get a new farthest during the current window,
// we don't need to update the cost
// as in a certain window introduced by an index, the cost of them are the same

package jump2

import "fmt"

func jump2(nums []int) int {
	n := len(nums)

	cost, curWindowRight, farthest := 0, 0, 0

	for i := 0; i < n-1; i++ {
		// i ~ i+nums[i] tells us the window.
		// during that window, the costs are the same
		// UNLESS, we see an index that was discovered by a previous window

		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == curWindowRight {
			cost += 1
			curWindowRight = farthest
			if curWindowRight >= n {
				return cost
			}
		}
	}

	return cost
}

func test() {
	/*
		Input: nums = [2,3,1,4,5]
		Output: 2
		Explanation: The minimum number of jumps
		to reach the last index is 2.
		Jump 1 step from index 0 to 1,
		then 3 steps to the last index.
	*/
	fmt.Println(jump2([]int{
		2, 3, 1, 4, 5,
	}))
}
