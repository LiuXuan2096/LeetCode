package Go_LeetCode

import "container/list"

func bstFromPreorder(preorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	N := len(preorder)
	nearBig := make([]int, N)
	for i := range nearBig {
		nearBig[i] = -1
	}
	stack := list.New()
	for i := 0; i < N; i++ {
		for stack.Len() > 0 && preorder[stack.Back().Value.(int)] < preorder[i] {
			nearBigIndex := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			nearBig[nearBigIndex] = i
		}
		stack.PushBack(i)
	}
	return process_1008(preorder, 0, N-1, nearBig)
}

func process_1008(pre []int, L, R int, nearBig []int) *TreeNode {
	if L > R {
		return nil
	}
	firstBig := R + 1
	if nearBig[L] != -1 && nearBig[L] <= R {
		firstBig = nearBig[L]
	}
	head := &TreeNode{Val: pre[L]}
	head.Left = process_1008(pre, L+1, firstBig-1, nearBig)
	head.Right = process_1008(pre, firstBig, R, nearBig)
	return head
}
