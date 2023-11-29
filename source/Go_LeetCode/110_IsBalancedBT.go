package Go_LeetCode

//func isBalanced(root *TreeNode) bool {
//	return process(root).isBalanced
//}
//
//type Info struct {
//	isBalanced bool
//	height     int
//}
//
//func process(node *TreeNode) Info {
//	if node == nil {
//		return Info{true, 0}
//	}
//	var leftInfo = process(node.Left)
//	var rightInfo = process(node.Right)
//	var height int = int(math.Max(float64(leftInfo.height), float64(rightInfo.height)) + 1)
//	var flag = true
//	if !leftInfo.isBalanced {
//		flag = false
//	}
//	if !rightInfo.isBalanced {
//		flag = false
//	}
//	if math.Abs(float64(leftInfo.height-rightInfo.height)) > 1 {
//		flag = false
//	}
//	return Info{flag, height}
//}
