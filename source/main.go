package main

import (
	"LeetCode-go/source/Go_LeetCode"
	"fmt"
)

func main() {
	arr := []int{3, 10, 5, 25, 2, 8}
	fmt.Println("Max XOR subarray:", Go_LeetCode.MaxXorSubarray2(arr))
}
