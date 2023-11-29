/*
https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
非递归实现二叉树的前序遍历
*/

package Go_LeetCode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0) // 用于存放前序遍历的结果
	if root != nil {
		stack := make([]*TreeNode, 0) // 使用该栈实现非递归前序遍历
		var ptr_node *TreeNode        // 用于遍历
		stack = append(stack, root)
		for len(stack) != 0 {
			ptr_node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 上方这两行代码放一起 才是完整的出栈操作
			res = append(res, (*ptr_node).Val)

			if (*ptr_node).Right != nil {
				stack = append(stack, (*ptr_node).Right) // 右子树入栈
			}
			if (*ptr_node).Left != nil {
				stack = append(stack, (*ptr_node).Left) // 左子树入栈
			}

		}
	}

	return res
}
