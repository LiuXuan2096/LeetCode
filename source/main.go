package main

import (
	"LeetCode-go/source/Go_LeetCode"
	"fmt"
)

func main() {
	var input [][]int
	var p []int = []int{1, 0}
	var p2 = []int{0, 1}
	input = append(input, p)
	input = append(input, p2)

	fmt.Println(Go_LeetCode.CanFinish(2, input))
}
