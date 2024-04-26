package Go_LeetCode

//import "fmt"
//
//func dfs(start *Node) {
//	if start == nil {
//		return
//	}
//	var stack []*Node
//	var sets map[*Node]void
//	var member void
//	stack = append(stack, start)
//	fmt.Println(start.value)
//	for len(stack) > 0 {
//		var cur *Node = stack[len(stack)-1]
//		stack = stack[0 : len(stack)-1]
//		for _, node := range cur.nexts {
//			_, ok := sets[node]
//			if !ok {
//				stack = append(stack, cur)
//				stack = append(stack, node)
//				sets[node] = member
//				fmt.Println(node.value)
//				break
//			}
//		}
//	}
//}
