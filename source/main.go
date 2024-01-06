package main

import (
	"LeetCode-go/source/Go_LeetCode"
	"fmt"
)

func main() {
	var nums = []int{3, 2, 1, 5, 6, 4, 6}
	var k = 2
	result := Go_LeetCode.FfindKthLargest(nums, k)
	fmt.Println(result)
}
