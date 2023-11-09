package Go_src

/**
* Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	n1 := head
	n2 := head
	for n2.Next != nil && n2.Next.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	// n1为链表中点

	n2 = n1.Next // n2是右边链表的头结点，利用n2将右半边链表倒序
	n1.Next = nil
	var n3 *ListNode
	for n2 != nil {
		n3 = n2.Next
		n2.Next = n1
		n1 = n2
		n2 = n3
	}
	n3 = n1   // n3保存了原链表的最后一个节点
	n2 = head // n2是链表左半边的第一个节点
	res := true
	for n2 != nil && n1 != nil {
		if n1.Val != n2.Val {
			res = false
			break
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	n1 = n3.Next
	n3.Next = nil
	for n1 != nil {
		n2 = n1.Next
		n1.Next = n3
		n3 = n1
		n1 = n2
	}
	return res

}
