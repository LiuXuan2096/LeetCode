package Go_LeetCode

import (
	"container/list"
)

/*
给定三个参数，二叉树的头节点head，树上某个节点target，
正数K。从target开始，
可以向上走或者向下走，返回与target的距离是K的所有节点
*/

type Node_03_08 struct {
	Value int
	Left  *Node_03_08
	Right *Node_03_08
}

func DistanceKNodes(root, target *Node_03_08, K int) []*Node_03_08 {
	parents := make(map[*Node_03_08]*Node_03_08)
	parents[root] = nil
	createParentMap(root, parents)
	queue := list.New()
	visited := make(map[*Node_03_08]bool)
	queue.PushBack(target)
	visited[target] = true
	curLevel := 0
	ans := make([]*Node_03_08, 0)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			cur := queue.Remove(queue.Front()).(*Node_03_08)
			if curLevel == K {
				ans = append(ans, cur)
			}
			if cur.Left != nil && !visited[cur.Left] {
				visited[cur.Left] = true
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil && !visited[cur.Right] {
				visited[cur.Right] = true
				queue.PushBack(cur.Right)
			}
			if parents[cur] != nil && !visited[parents[cur]] {
				visited[parents[cur]] = true
				queue.PushBack(parents[cur])
			}
		}
		curLevel++
		if curLevel > K {
			break
		}
	}
	return ans
}

func createParentMap(cur *Node_03_08, parents map[*Node_03_08]*Node_03_08) {
	if cur == nil {
		return
	}
	if cur.Left != nil {
		parents[cur.Left] = cur
		createParentMap(cur.Left, parents)
	}
	if cur.Right != nil {
		parents[cur.Right] = cur
		createParentMap(cur.Right, parents)
	}
}

/*
测试代码：
func main() {
	n0 := &Go_LeetCode.Node_03_08{Value: 0}
	n1 := &Go_LeetCode.Node_03_08{Value: 1}
	n2 := &Go_LeetCode.Node_03_08{Value: 2}
	n3 := &Go_LeetCode.Node_03_08{Value: 3}
	n4 := &Go_LeetCode.Node_03_08{Value: 4}
	n5 := &Go_LeetCode.Node_03_08{Value: 5}
	n6 := &Go_LeetCode.Node_03_08{Value: 6}
	n7 := &Go_LeetCode.Node_03_08{Value: 7}
	n8 := &Go_LeetCode.Node_03_08{Value: 8}

	n3.Left = n5
	n3.Right = n1
	n5.Left = n6
	n5.Right = n2
	n1.Left = n0
	n1.Right = n8
	n2.Left = n7
	n2.Right = n4

	root := n3
	target := n5
	K := 2

	ans := Go_LeetCode.DistanceKNodes(root, target, K)
	for _, Node_03_08 := range ans {
		fmt.Println(Node_03_08.Value)
	}
}
*/
