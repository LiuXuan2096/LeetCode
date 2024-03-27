package Go_LeetCode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		} else {
			return list1
		}
	}
	var head, cur1, cur2, pre *ListNode
	if list1.Val <= list2.Val {
		head = list1
		cur1 = list1.Next
		cur2 = list2
	} else {
		head = list2
		cur1 = list1
		cur2 = list2.Next
	}
	pre = head
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}
	if cur1 != nil {
		pre.Next = cur1
	} else {
		pre.Next = cur2
	}
	return head
}
