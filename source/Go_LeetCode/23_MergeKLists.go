package Go_LeetCode

import "container/heap"

type ListNodeHeap []*ListNode

/*
Len Less Swap Push Pop
*/
func (h ListNodeHeap) Len() int {
	return len(h)
}

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	h := &ListNodeHeap{}
	heap.Init(h)
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}
	if h.Len() == 0 {
		return nil
	}
	head := heap.Pop(h).(*ListNode)
	pre := head
	if pre.Next != nil {
		heap.Push(h, pre.Next)
	}
	for h.Len() > 0 {
		cur := heap.Pop(h).(*ListNode)
		pre.Next = cur
		pre = cur
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}
	return head
}
