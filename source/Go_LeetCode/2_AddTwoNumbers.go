package Go_LeetCode

//type ListNode struct {
//	     Val int
//	     Next *ListNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0 // 表示进位
	n1, n2, n := 0, 0, 0
	curNode1, curNode2 := l1, l2
	var node, pre *ListNode

	for curNode1 != nil || curNode2 != nil {
		n1 = 0
		n2 = 0
		if curNode1 != nil {
			n1 = curNode1.Val
			curNode1 = curNode1.Next
		}
		if curNode2 != nil {
			n2 = curNode2.Val
			curNode2 = curNode2.Next
		}

		n = n1 + n2 + carry
		pre = node
		node = &ListNode{Val: n % 10, Next: pre}
		carry = n / 10
	}

	if carry == 1 {
		pre = node
		node = &ListNode{Val: 1, Next: pre}
	}

	return reverseList(node)
}

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	var curNode = head
	for curNode != nil {
		next = curNode.Next
		curNode.Next = pre
		pre = curNode
		curNode = next
	}
	return pre
}
