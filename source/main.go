package main

import (
	"LeetCode-go/source/Go_src"
	"fmt"
)

func main() {
	node3 := Go_src.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node2 := Go_src.TreeNode{
		Val:   2,
		Left:  &node3,
		Right: nil,
	}
	node1 := Go_src.TreeNode{
		Val:   1,
		Left:  nil,
		Right: &node2,
	}
	res := Go_src.InorderTraversal(&node1)
	fmt.Println(res)
}
