package Go_LeetCode

/*
程序有问题，
*/

import (
	"fmt"
	"math"
)

type Node1 struct {
	happyVal int
	children []*Node1
}

func Main1() {
	number := 0
	root := 0
	happy := 0
	_, _ = fmt.Scan(&number, &root)

	staff := make([]*Node1, number)

	for i := 0; i < number; i++ {
		_, _ = fmt.Scan(&happy)
		var tempNode Node1 = Node1{happy, nil}
		staff = append(staff, &tempNode)
	}

	for {
		var a, b int
		_, _ = fmt.Scan(&a, &b)
		staff[a].children = append(staff[a].children, staff[b])
	}

	var allInfo *Info1 = process1(staff[root])
	fmt.Println(math.Max(float64(allInfo.no), float64(allInfo.yes)))
}

type Info1 struct {
	no  int // 该节点没来时，以该节点为根节点的树的最大快乐值
	yes int // 该节点来时，以该节点为根节点的树的最大快乐值
}

//func process1(node *Node1) *Info1 {
//	if node == nil {
//		return &Info1{0, 0}
//	}
//	var no = 0
//	var yes = node.happyVal
//	for _, val := range node.children {
//		var info *Info1 = process1(val)
//		no += int(math.Max(float64(info.no), float64(info.yes)))
//		yes += info.yes
//	}
//	return &Info1{no, yes}
//}
