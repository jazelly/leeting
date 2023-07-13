package distanceK

// Beats 100% time and 50% space

func distanceK2(root *TreeNode, target *TreeNode, k int) []int {
	var (
		result  []int
		parents = make(map[int]*TreeNode)
		visit   = make(map[int]bool)
		travel  func(*TreeNode)
		dfs     func(*TreeNode, int)
	)

	travel = func(t *TreeNode) {
		if t.Left != nil {
			parents[t.Left.Val] = t
			travel(t.Left)
		}
		if t.Right != nil {
			parents[t.Right.Val] = t
			travel(t.Right)
		}
	}
	travel(root)

	dfs = func(t *TreeNode, d int) {
		if t == nil || d > k || visit[t.Val] {
			return
		}
		visit[t.Val] = true
		if d == k {
			result = append(result, t.Val)
			return
		}
		dfs(t.Left, d+1)
		dfs(t.Right, d+1)
		dfs(parents[t.Val], d+1)
	}
	dfs(target, 0)

	return result
}
