package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		s1, c1 := dfs(node.Left)
		s2, c2 := dfs(node.Right)

		tc := c1 + c2 + 1
		ts := s1 + node.Val + s2
		ta := ts / tc

		if ta == node.Val {
			res++
		}

		return ts, tc
	}
	dfs(root)
	return res
}

func main() {
	root := &TreeNode{Val: 4}
	eight := &TreeNode{Val: 8}
	five := &TreeNode{Val: 5}
	o := &TreeNode{Val: 0}
	one := &TreeNode{Val: 1}
	six := &TreeNode{Val: 6}
	root.Left = eight
	root.Right = five
	eight.Left = o
	eight.Right = one
	five.Right = six

	averageOfSubtree(root)
}
