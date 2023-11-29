/*
https://leetcode.cn/problems/binary-tree-inorder-traversal/
非递归方式实现二叉树的中序遍历
*/

package Go_LeetCode

func InorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil {
		stack := make([]*TreeNode, 0)
		var current = root
		for len(stack) != 0 || current != nil {
			if current != nil {
				stack = append(stack, current) //当前节点入栈
				current = (*current).Left
			} else {
				// 当前节点为空，访问栈顶元素,并让栈顶元素出栈
				current = stack[len(stack)-1]
				stack = stack[:len(stack)-1]      // pop
				res = append(res, (*current).Val) // 访问出栈的元素
				current = (*current).Right        // 访问右子树

			}
		}
	}

	return res
}
