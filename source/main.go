package main

import (
	"LeetCode-go/source/Go_LeetCode"
	"fmt"
)

func main() {
	//GSY11.GSY()
	//var nums = []int{3, 2, 1, 5, 6, 4, 6}
	//var k = 2
	//result := Go_LeetCode.FfindKthLargest(nums, k)
	//fmt.Println(result)

	ac := Go_LeetCode.NewACAutomation()
	ac.Insert("dhe")
	ac.Insert("he")
	ac.Insert("abcdheks")
	ac.Build()

	contains := ac.ContainWords("abcdhekskdjfafhasldkflskdjhwqaeruv")
	for _, word := range contains {
		fmt.Println(word)
	}
}
