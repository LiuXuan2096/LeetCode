package Go_LeetCode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return process(root, p, q).ans
}

type Info struct {
	findA bool
	findB bool
	ans   *TreeNode
}

func process(node *TreeNode, p *TreeNode, q *TreeNode) *Info {
	if node == nil {
		return &Info{false, false, nil}
	}
	var leftInfo = process(node.Left, p, q)
	var rightInfo = process(node.Right, p, q)

	var findA bool = (node == p) || leftInfo.findA || rightInfo.findA
	var findB bool = (node == q) || leftInfo.findB || rightInfo.findB

	var ans *TreeNode = nil
	if leftInfo.ans != nil {
		ans = leftInfo.ans
	} else if rightInfo.ans != nil {
		ans = rightInfo.ans
	} else if findA && findB {
		ans = node
	}

	return &Info{findA, findB, ans}
}
