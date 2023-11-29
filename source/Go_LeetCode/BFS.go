package Go_LeetCode

import "fmt"

type Node struct {
	value int
	in    int
	out   int
	nexts []*Node
	edges []*Edge
}

type Edge struct {
	weight int
	from   Node
	to     Node
}

type void struct {
}

func bfs(start *Node) {
	if start == nil {
		return
	}
	var queue []*Node
	var set map[*Node]void
	var member void
	queue = append(queue, start)
	set[start] = member

	for len(queue) > 0 {
		var tempNode *Node = queue[0]
		queue = queue[1:]
		fmt.Println(tempNode.value)
		for _, val := range tempNode.nexts {
			_, ok := set[val]
			if !ok {
				queue = append(queue, val)
				set[val] = member
			}
		}
	}
}
