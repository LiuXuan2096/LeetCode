package Go_LeetCode

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return p(root)
}

// 返回x为头的树， 最小深度是多少
func p(x *TreeNode) int {
	if x.Left == nil && x.Right == nil {
		return 1
	}
	// 左右子树起码一个不为空
	var leftH = math.MaxInt32
	if x.Left != nil {
		leftH = p(x.Left)
	}
	var rightH = math.MaxInt32
	if x.Right != nil {
		rightH = p(x.Right)
	}
	return int(1 + math.Min(float64(leftH), float64(rightH)))
}
