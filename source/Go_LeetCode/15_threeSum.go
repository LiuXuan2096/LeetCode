package Go_LeetCode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}

		}
	}
	return res
}

// threeSum1 双指针解法
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			nexts := twoSum1(nums, i+1, -nums[i])
			for _, cur := range nexts {
				cur = append([]int{nums[i]}, cur...)
				ans = append(ans, cur)
			}
		}
	}
	return ans
}

// twoSum1 此方法假定传进来的nums是有序的，
// 传入 nums 前需要对nums进行排序
func twoSum1(nums []int, begin int, target int) [][]int {
	L := begin
	R := len(nums) - 1
	ans := [][]int{}
	for L < R {
		if nums[L]+nums[R] > target {
			R--
		} else if nums[L]+nums[R] < target {
			L++
		} else {
			if L == begin || nums[L-1] != nums[L] {
				ans = append(ans, []int{nums[L], nums[R]})
			}
			L++
		}
	}
	return ans
}
