package Go_src

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	var curEnd *TreeNode = root // 当前层，最右节点是谁
	var nextEnd *TreeNode = nil // 下一层，最右节点是谁
	s := []int{root.Val}
	ans = append(ans, s)
	var levelNodes []int

	for len(queue) != 0 {
		var cur *TreeNode = queue[0]
		queue = queue[1:]

		if cur.Left != nil {
			queue = append(queue, cur.Left)
			nextEnd = cur.Left
			levelNodes = append(levelNodes, cur.Left.Val)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			nextEnd = cur.Right
			levelNodes = append(levelNodes, cur.Right.Val)
		}
		if cur == curEnd {
			curEnd = nextEnd
			if len(levelNodes) != 0 {
				ans = append(ans, levelNodes)
			}
			levelNodes = []int{}
		}
	}
	return ans
}
