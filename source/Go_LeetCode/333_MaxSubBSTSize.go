package Go_LeetCode

//func largestBSTSubtree(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return process(root).maxBSTSubTreeSize
//}
//
//type Info struct {
//	maxBSTSubTreeSize int
//	allSize           int
//	max               int
//	min               int
//}
//
//func process(node *TreeNode) *Info {
//	if node == nil {
//		return nil
//	}
//	var leftInfo = process(node.Left)
//	var rightInfo = process(node.Right)
//	var max int = node.Val
//	var min int = node.Val
//	var allSize int = 1
//
//	if leftInfo != nil {
//		max = int(math.Max(float64(leftInfo.max), float64(max)))
//		min = int(math.Min(float64(leftInfo.min), float64(min)))
//		allSize += leftInfo.allSize
//	}
//
//	if rightInfo != nil {
//		max = int(math.Max(float64(rightInfo.max), float64(max)))
//		min = int(math.Min(float64(rightInfo.min), float64(min)))
//		allSize += rightInfo.allSize
//	}
//
//	var case1 int = -1 // 左子树的最大二叉搜索子数的节点数
//	if leftInfo != nil {
//		case1 = leftInfo.maxBSTSubTreeSize
//	}
//
//	var case2 int = -1 // 右子树的最大二叉搜索子树的节点数
//	if rightInfo != nil {
//		case2 = rightInfo.maxBSTSubTreeSize
//	}
//
//	var case3 int = -1 // 如果该节点的左右子树都是二叉搜索树的话，返回以该节点为跟节点的整棵树的节点数
//	var leftBST bool
//	var rightBST bool
//	if leftInfo == nil || leftInfo.maxBSTSubTreeSize == leftInfo.allSize {
//		leftBST = true
//	} else {
//		leftBST = false
//	}
//	if rightInfo == nil || rightInfo.maxBSTSubTreeSize == rightInfo.allSize {
//		rightBST = true
//	} else {
//		rightBST = false
//	}
//	if leftBST && rightBST {
//		var leftMaxLessParent bool
//		var rightMinMoreParent bool
//		if leftInfo == nil || leftInfo.max < node.Val {
//			leftMaxLessParent = true
//		} else {
//			leftMaxLessParent = false
//		}
//		if rightInfo == nil || rightInfo.min > node.Val {
//			rightMinMoreParent = true
//		} else {
//			rightMinMoreParent = false
//		}
//
//		if leftMaxLessParent && rightMinMoreParent {
//			var leftSize int
//			var rightSize int
//			if leftInfo == nil {
//				leftSize = 0
//			} else {
//				leftSize = leftInfo.allSize
//			}
//			if rightInfo == nil {
//				rightSize = 0
//			} else {
//				rightSize = rightInfo.allSize
//			}
//			case3 = leftSize + rightSize + 1
//		}
//	}
//
//	return &Info{int(math.Max(float64(case1), math.Max(float64(case2), float64(case3)))), allSize, max, min}
//
//}
