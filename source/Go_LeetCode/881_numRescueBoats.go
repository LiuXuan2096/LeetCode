package Go_LeetCode

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	ans := 0
	l, r := 0, len(people)-1
	for l <= r {
		sum := 0
		if l == r {
			sum = people[l]
		} else {
			sum = people[l] + people[r]
		}
		if sum > limit {
			r--
		} else {
			l++
			r--
		}
		ans++
	}
	return ans
}
