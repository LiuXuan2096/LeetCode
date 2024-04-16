package Go_LeetCode

//
//import (
//	"crypto/sha256"
//	"encoding/hex"
//)
//
//type Info struct {
//	ans int
//	str string
//}
//
//func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
//	ans := make([]*TreeNode, 0)
//	nodeSet := make(map[*TreeNode]struct{})
//	_ = process_652(root, &ans, nodeSet)
//	return ans
//}
//
//func process_652(head *TreeNode, p_ans *[]*TreeNode, set map[*TreeNode]struct{}) Info {
//	ans := *p_ans
//	if head.Left == nil && head.Right == nil {
//		// 叶子节点
//		if _, ok := set[head]; !ok {
//			ans = append(ans, head)
//			set[head] = struct{}{}
//		}
//		return Info{ans: 0, str: getHashCode("#,")}
//	} else {
//		var l, r Info
//		if head.Left != nil {
//			l = process_652(head.Left, p_ans, set)
//		}
//		if head.Right != nil {
//			r = process_652(head.Right, p_ans, set)
//		}
//		if l.str == r.str {
//			if _, ok := set[head]; !ok {
//				ans = append(ans, head)
//				set[head]
//			}
//		}
//	}
//
//}
//
//func getHashCode(input string) string {
//	hasher := sha256.New()
//	hasher.Write([]byte(input))
//	return hex.EncodeToString(hasher.Sum(nil))
//}
