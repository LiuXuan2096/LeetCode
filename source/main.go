package main

import (
	"LeetCode-go/source/Go_LeetCode"
	"fmt"
)

func main() {
	arr := []int{3, 1, 4, 1, 5}
	K := 5
	fmt.Println(Go_LeetCode.GetMaxLessOrEqualK(arr, K))
}
