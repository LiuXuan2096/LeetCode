package Go_LeetCode

func recoverTree(root *TreeNode) {
	x := (*TreeNode)(nil) // 使用指针来记录需要交换的节点x
	y := (*TreeNode)(nil) // 使用指针来记录需要交换的节点y
	pre := (*TreeNode)(nil)

	dfs(root, &x, &y, &pre)
	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}
}

// dfs 用于中序遍历二叉树，并比较上一个节点(pre)和当前节点的值
func dfs(node *TreeNode, x, y, pre **TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left, x, y, pre)
	if *pre == nil {
		*pre = node
	} else {
		if (*pre).Val > node.Val {
			*y = node
			if *x == nil {
				*x = *pre
			}
		}
		*pre = node
	}
	dfs(node.Right, x, y, pre)
}
