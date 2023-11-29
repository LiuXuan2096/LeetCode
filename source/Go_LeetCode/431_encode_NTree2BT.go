package Go_LeetCode

//
//type Node struct {
//	Val      int
//	Children []*Node
//}
//
//type Codec struct {
//}
//
//func Constructor() *Codec {
//	return new(Codec)
//}
//
//func (this *Codec) encode(root *Node) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	var head TreeNode = TreeNode{root.Val, nil, nil}
//	head.Left = en(root.Children)
//	return &head
//}
//
//func en(children []*Node) *TreeNode {
//	var head *TreeNode = nil
//	var cur *TreeNode = nil
//	for _, val := range children {
//		var ptrTempNode *TreeNode = &TreeNode{val.Val, nil, nil}
//		if head == nil {
//			head = ptrTempNode
//		} else {
//			cur.Right = ptrTempNode
//		}
//		cur = ptrTempNode
//		cur.Left = en(val.Children)
//	}
//	return head
//}
//
//func (this *Codec) decode(root *TreeNode) *Node {
//	if root == nil {
//		return nil
//	}
//	return &Node{root.Val, de(root.Left)}
//}
//
//func de(root *TreeNode) []*Node {
//	var children []*Node
//	for root != nil {
//		var cur *Node = &Node{root.Val, de(root.Left)}
//		children = append(children, cur)
//		root = root.Right
//	}
//	return children
//}
