package main

import (
	"LeetCode-go/source/from_forceRecursive_to_DP"
	"fmt"
)

func main() {
	arr := []int{5, 7, 4, 5, 8, 1, 6, 0, 3, 4, 6, 1, 7}
	result := from_forceRecursive_to_DP.Win3(arr)
	fmt.Printf("Winning Score: %d\n", result)
}
