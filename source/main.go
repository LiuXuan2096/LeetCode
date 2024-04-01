package main

import (
	"LeetCode-go/source/Go_LeetCode"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := 11
	fmt.Println(Go_LeetCode.NumSubArray_滑动窗口(arr, num))
}
