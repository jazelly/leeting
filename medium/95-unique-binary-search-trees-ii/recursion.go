package UniqueBinarySearchTrees2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	memo := make(map[[2]int][]*TreeNode)

	var dfs func(s int, e int) []*TreeNode
	dfs = func(s int, e int) []*TreeNode {
		thisRange := [2]int{s, e}
		v, exists := memo[thisRange]
		if exists {
			return v
		}

		buffer := []*TreeNode{}
		if s > e {
			buffer = append(buffer, nil)
			return buffer
		}

		for i := s; i <= e; i++ {
			l := dfs(s, i-1)
			r := dfs(i+1, e)
			for _, lnode := range l {
				for _, rnode := range r {
					root := &TreeNode{
						Val:   i,
						Left:  lnode,
						Right: rnode,
					}
					buffer = append(buffer, root)
				}
			}
		}
		memo[thisRange] = buffer
		return buffer
	}

	return dfs(1, n)
}
