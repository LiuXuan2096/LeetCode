package Go_LeetCode

//func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return process(root).isBST
//}
//
//type Info struct {
//	isBST bool
//	max   int
//	min   int
//}
//
//func process(node *TreeNode) *Info {
//	if node == nil {
//		return nil
//	}
//	var leftInfo = process(node.Left)
//	var rightInfo = process(node.Right)
//
//	var max int = node.Val
//	if leftInfo != nil {
//		max = int(math.Max(float64(max), float64(leftInfo.max)))
//	}
//	if rightInfo != nil {
//		max = int(math.Max(float64(max), float64(rightInfo.max)))
//	}
//
//	var min int = node.Val
//	if leftInfo != nil {
//		min = int(math.Min(float64(min), float64(leftInfo.min)))
//	}
//	if rightInfo != nil {
//		min = int(math.Min(float64(min), float64(rightInfo.min)))
//	}
//
//	var isBst bool = true
//	if leftInfo != nil && !leftInfo.isBST {
//		isBst = false
//	}
//	if rightInfo != nil && !rightInfo.isBST {
//		isBst = false
//	}
//	if leftInfo != nil && leftInfo.max >= node.Val {
//		isBst = false
//	}
//	if rightInfo != nil && rightInfo.min <= node.Val {
//		isBst = false
//	}
//	return &Info{isBst, max, min}
//}
