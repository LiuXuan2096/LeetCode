package Go_LeetCode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		n--
		if n == -1 {
			pre = head
		}
		if n < -1 {
			pre = pre.Next
		}
		cur = cur.Next
	}
	if n > 0 {
		return head
	}
	if pre == nil {
		return head.Next
	}
	pre.Next = pre.Next.Next
	return head
}
