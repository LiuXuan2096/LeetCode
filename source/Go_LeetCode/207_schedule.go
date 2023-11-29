package Go_LeetCode

type Node207 struct {
	value int
	in    int // 该节点入度
	out   int // 该节点出度
	nexts []*Node207
}

var member void

func CanFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 1 || len(prerequisites) == 0 {
		return true
	}

	var ans = false
	nodes := make(map[int]*Node207, 10)
	var zeroInQueue []*Node207
	var help207 []*Node207

	for _, pre_after := range prerequisites {
		var pre_node *Node207
		var after_node *Node207
		_, ok := nodes[pre_after[0]]
		if !ok {
			// 该节点未创建
			var tempNode = &Node207{value: pre_after[0], in: 0, out: 1}
			nodes[pre_after[0]] = tempNode
			pre_node = tempNode
		} else {
			// 该节点已创建
			nodes[pre_after[0]].out++
			pre_node = nodes[pre_after[0]]
		}

		_, ok = nodes[pre_after[1]]
		if !ok {
			// 该节点未创建
			var tempNode = &Node207{value: pre_after[1], in: 1, out: 0}
			nodes[pre_after[1]] = tempNode
			after_node = tempNode
		} else {
			// 该节点已创建
			nodes[pre_after[1]].in++
			after_node = nodes[pre_after[1]]
		}
		pre_node.nexts = append(pre_node.nexts, after_node)
	}

	for _, val := range nodes {
		if val.in == 0 {
			zeroInQueue = append(zeroInQueue, val)
			help207 = append(help207, val)
		}
	}

	if len(zeroInQueue) == 0 {
		ans = false
	}

	for len(help207) > 0 {
		var cur = help207[0]
		help207 = help207[1:]
		for _, n := range cur.nexts {
			n.in--
			if n.in == 0 {
				zeroInQueue = append(zeroInQueue, n)
				help207 = append(help207, n)
			}
		}
	}

	if len(zeroInQueue) == len(nodes) {
		ans = true
	}

	return ans
}
