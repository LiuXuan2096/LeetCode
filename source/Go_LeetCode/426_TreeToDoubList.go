package Go_LeetCode

type Info_426 struct {
	start *Node
	end   *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	allInfo := process_426(root)
	allInfo.end.Right = allInfo.start
	allInfo.start.Left = allInfo.end
	return allInfo.start
}

func process_426(X *Node) Info_426 {
	if X == nil {
		return Info_426{nil, nil}
	}
	lInfo := process_426(X.Left)
	rInfo := process_426(X.Right)
	if lInfo.end != nil {
		lInfo.end.Right = X
	}
	X.Left = lInfo.end
	X.Right = rInfo.start
	if rInfo.start != nil {
		rInfo.start.Left = X
	}
	// 整体链表的头    lInfo.start != null ? lInfo.start : X
	// 整体链表的尾    rInfo.end != null ? rInfo.end : X
	var start, end = X, X
	if lInfo.start != nil {
		start = lInfo.start
	}
	if rInfo.end != nil {
		end = rInfo.end
	}
	return Info_426{start, end}
}
