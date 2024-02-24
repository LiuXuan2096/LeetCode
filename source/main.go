package main

import (
	"LeetCode-go/source/from_forceRecursive_to_DP"
	"fmt"
)

func main() {
	s1 := "a12b3c456d"
	s2 := "1ef23ghi4j56k"
	result := from_forceRecursive_to_DP.LongestCommonSubsequence2(s1, s2)
	fmt.Println(result)
}
