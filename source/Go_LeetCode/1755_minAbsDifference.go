package Go_LeetCode

import (
	"math"
	"sort"
)

// 本题测试链接 : https://leetcode.com/problems/closest-subsequence-sum/
// 本题数据量描述:
// 1 <= nums.length <= 40
// -10^7 <= nums[i] <= 10^7
// -10^9 <= goal <= 10^9
// 通过这个数据量描述可知，需要用到分治，因为数组长度不大
// 而值很大，用动态规划的话，表会爆
func minAbsDifference(nums []int, goal int) int {
	if nums == nil || len(nums) == 0 {
		return goal
	}
	// 数组左半部分的所有可能累加和都放在这里
	l := make([]int, 1<<20)
	// 数组右半部分的所有可能累加和都放在这里
	r := make([]int, 1<<20)

	// 收集数组左半部分，所有可能的累加和
	// 并且返回一共收集了几个数，就是le
	le := process_1755(nums, 0, len(nums)/2, 0, 0, l)
	// 收集数组右半部分，所有可能的累加和
	// 并且返回一共收集了几个数，就是re
	re := process_1755(nums, len(nums)/2, len(nums), 0, 0, r)
	// 把左半部分收集到的累加和排序
	sort.Ints(l[:le])
	// 把右半部分收集到的累加和排序
	sort.Ints(r[:re])
	re--
	ans := math.Abs(float64(goal))
	// 为什么要排序？
	// 因为排序之后定位数字可以不回退
	// 比如你想累加和尽量接近10
	// 左半部分累加和假设为 : 0, 2, 3, 7, 9
	// 右半部分累加和假设为 : 0, 6, 7, 8, 9
	// 左部分累加和是0时，右半部分选哪个累加和最接近10，是9
	// 左部分累加和是2时，右半部分选哪个累加和最接近10，是8
	// 左部分累加和是3时，右半部分选哪个累加和最接近10，是7
	// 左部分累加和是7时，右半部分选哪个累加和最接近10，是0
	// 左部分累加和是9时，右半部分选哪个累加和最接近10，是0
	// 上面你可以看到，
	// 当你从左往右选择左部分累加和时，右部分累加和的选取，可以从右往左
	// 这就非常的方便
	// 下面的代码就是这个意思
	for i := 0; i < le; i++ {
		rest := goal - l[i]
		for re > 0 && math.Abs(float64(rest-r[re-1])) <= math.Abs(float64(rest-r[re])) {
			re--
		}
		ans = math.Min(ans, math.Abs(float64(rest-r[re])))
	}
	return int(ans)
}

// 解释一下
// nums[0..index-1]已经选了一些数字，组成了累加和sum
// 当前来到nums[index....end)这个范围，所有可能的累加和
// 填写到arr里去
// fill参数的意思是: 如果出现新的累加和，填写到arr的什么位置
// 返回所有生成的累加和，现在填到了arr的什么位置
func process_1755(nums []int, i, end, sum, fill int, arr []int) int {
	if i == end {
		// 到了终止为止了，该结束了
		// 把当前的累加和sum
		// 填写到arr[fill]的位置
		// 然后fill++，表示如果后续再填的话
		// 该放在什么位置了
		arr[fill] = sum
		fill++
	} else {
		// 可能性1 : 不要当前的数字
		// 走一个分支，形成多少累加和，都填写到arr里去
		// 同时返回这个分支把arr填到了什么位置
		fill = process_1755(nums, i+1, end, sum, fill, arr)
		// 可能性2 : 要当前的数字
		// 走一个分支，形成多少累加和，都填写到arr里去
		// 接着可能性1所填到的位置，继续填写到arr里去
		// 这就是为什么要拿到上一个分支填到哪了
		// 因为如果没有这个信息，可能性2的分支不知道往哪填生成的累加和
		fill = process_1755(nums, i+1, end, sum+nums[i], fill, arr)
	}
	return fill
}
