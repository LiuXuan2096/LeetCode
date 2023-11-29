package Go_LeetCode

//func isCompleteTree(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return process(root).isCBT
//}
//
//type Info struct {
//	isFull bool
//	isCBT  bool
//	height int
//}
//
//func process(node *TreeNode) *Info {
//	if node == nil {
//		return &Info{true, true, 0}
//	}
//
//	var leftInfo = process(node.Left)
//	var rightInfo = process(node.Right)
//
//	var height int = int(math.Max(float64(leftInfo.height), float64(rightInfo.height)) + 1)
//	var isFull bool = leftInfo.isFull && rightInfo.isFull && leftInfo.height == rightInfo.height
//
//	var isCBT = false
//	if isFull {
//		isCBT = true
//	} else {
//		// 这个分支说明以node节点为根节点的整棵树不满
//		if leftInfo.isCBT && rightInfo.isCBT {
//			if leftInfo.isCBT && rightInfo.isFull && leftInfo.height == rightInfo.height+1 {
//				isCBT = true
//			}
//			if leftInfo.isFull && rightInfo.isFull && leftInfo.height == rightInfo.height+1 {
//				isCBT = true
//			}
//			if leftInfo.isFull && rightInfo.isCBT && leftInfo.height == rightInfo.height {
//				isCBT = true
//			}
//		}
//	}
//
//	return &Info{isFull, isCBT, height}
//
//}
