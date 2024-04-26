package Go_LeetCode

import "math"

/*
 * 给定一个数组arr，长度为N，arr中的值不是0就是1
 * arr[i]表示第i栈灯的状态，0代表灭灯，1代表亮灯
 * 每一栈灯都有开关，但是按下i号灯的开关，会同时改变i-1、i、i+2栈灯的状态
 * 问题一：
 * 如果N栈灯排成一条直线,请问最少按下多少次开关,能让灯都亮起来
 * 排成一条直线说明：
 * i为中间位置时，i号灯的开关能影响i-1、i和i+1
 * 0号灯的开关只能影响0和1位置的灯
 * N-1号灯的开关只能影响N-2和N-1位置的灯
 *
 * 问题二：
 * 如果N栈灯排成一个圈,请问最少按下多少次开关,能让灯都亮起来
 * 排成一个圈说明：
 * i为中间位置时，i号灯的开关能影响i-1、i和i+1
 * 0号灯的开关能影响N-1、0和1位置的灯
 * N-1号灯的开关能影响N-2、N-1和0位置的灯
 *
 */

func NoLoopMinStep(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0] ^ 1
	}
	if len(arr) == 2 {
		if arr[0] != arr[1] {
			return math.MaxInt32
		}
		return arr[0] ^ 1
	}
	// 不变0位置的状态
	p1 := process_09_01(arr, 2, arr[0], arr[1])
	// 改变0位置的状态
	p2 := process_09_01(arr, 2, arr[0]^1, arr[1]^1)
	if p2 != math.MaxInt32 {
		p2++
	}
	return min_09_01(p1, p2)
}

// 当前在哪个位置上，做选择，nextIndex - 1 = cur ，当前！
// cur - 1 preStatus
// cur  curStatus
// 0....cur-2  全亮的！
func process_09_01(arr []int, nextIndex, preStatus, curStatus int) int {
	// 当前来到最后一个开关的位置
	if nextIndex == len(arr) {
		if preStatus != curStatus {
			return math.MaxInt32
		}
		return curStatus ^ 1
	}
	// 没到最后一个按钮呢！
	// i < arr.length
	if preStatus == 0 {
		// 当前位置的前一个位置灯还没开
		// 所以一定要改变前一个位置的灯的状态
		curStatus ^= 1            //表示按下当前位置的开关
		cur := arr[nextIndex] ^ 1 //当前位置开关按下后，下一个位置的状态也随之变化
		next := process_09_01(arr, nextIndex+1, curStatus, cur)
		if next == math.MaxInt32 {
			return next
		}
		return next + 1
	} else {
		// 前一个位置灯是开着的
		// 所以一定不能改变前一个位置灯的状态
		return process_09_01(arr, nextIndex+1, curStatus, arr[nextIndex])
	}
}

func min_09_01(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
