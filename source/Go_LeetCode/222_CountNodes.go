package Go_LeetCode

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return bs(root, 1, mostLeftLevel(root, 1))
}

// bs 根据当前节点、当前层级和总层级计算子树节点数
// 当前来到node节点，node节点在level层，总层数是h
// 返回node为头的子树(必是完全二叉树)，有多少个节点
func bs(node *TreeNode, level, h int) int {
	if level == h {
		return 1
	}
	if mostLeftLevel(node.Right, level+1) == h {
		// 当前节点右子树的最大深度=当前节点树的最大深度
		// 说明当前节点左子树一定是满二叉树 再继续计算右子树的节点个数
		return pow_222(h-level) + bs(node.Right, level+1, h)
	} else {
		// 当前节点右子树的最大深度=当前节点树的最大深度-1
		// 说明当前节点右子树一点是满二叉树 再继续计算左子树的节点数
		return pow_222(h-level-1) + bs(node.Left, level+1, h)
	}
}

func pow_222(a int) int {
	return int(math.Pow(2, float64(a)))
}

// mostLeftLevel 找出以node为头结点的子树的最大深度
// 如果node在第level层，
// 求以node为头的子树，最大深度是多少
// node为头的子树，一定是完全二叉树
func mostLeftLevel(node *TreeNode, level int) int {
	for node != nil {
		level++
		node = node.Left
	}
	return level - 1
}
