package Go_LeetCode

import (
	"container/list"
	"math"
)

func NnumSubmat(mat [][]int) int {
	if mat == nil || len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	var nums = 0
	var height = make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		nums += countFromBottom(height)
	}
	return nums
}

// 比如
//              1
//              1
//              1         1
//    1         1         1
//    1         1         1
//    1         1         1
//
//    2  ....   6   ....  9
// 如上图，假设在6位置，1的高度为6
// 在6位置的左边，离6位置最近、且小于高度6的位置是2，2位置的高度是3
// 在6位置的右边，离6位置最近、且小于高度6的位置是9，9位置的高度是4
// 此时我们求什么？
// 1) 求在3~8范围上，必须以高度6作为高的矩形，有几个？
// 2) 求在3~8范围上，必须以高度5作为高的矩形，有几个？
// 也就是说，<=4的高度，一律不求
// 那么，1) 求必须以位置6的高度6作为高的矩形，有几个？
// 3..3  3..4  3..5  3..6  3..7  3..8
// 4..4  4..5  4..6  4..7  4..8
// 5..5  5..6  5..7  5..8
// 6..6  6..7  6..8
// 7..7  7..8
// 8..8
// 这么多！= 21 = (9 - 2 - 1) * (9 - 2) / 2
// 这就是任何一个数字从栈里弹出的时候，计算矩形数量的方式
func countFromBottom(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	var nums = 0
	var stack list.List
	for i := 0; i < len(height); i++ {
		for stack.Len() > 0 && height[i] <= height[stack.Front().Value.(int)] {
			var cur = stack.Front().Value.(int)
			stack.Remove(stack.Front())
			if height[cur] > height[i] {
				var left = -1
				var down int // 表示 cur 位置的值左右两边比它小的两个值中 最大的那个值的下标
				if stack.Len() > 0 {
					left = stack.Front().Value.(int)
				}
				var n = i - left - 1
				if left == -1 {
					down = int(math.Max(0, float64(height[i])))
				} else {
					down = int(math.Max(float64(height[left]), float64(height[i])))
				}
				nums += (height[cur] - down) * num(n)
			}
		}
		stack.PushFront(i)
	}
	for stack.Len() > 0 {
		var cur = stack.Front().Value.(int)
		stack.Remove(stack.Front())
		left := -1
		if stack.Len() > 0 {
			left = stack.Front().Value.(int)
		}
		var n = len(height) - left - 1
		var down int
		if left == -1 {
			down = 0
		} else {
			down = height[left]
		}
		nums += (height[cur] - down) * num(n)
	}
	return nums
}

func num(n int) int {
	return (n * (n + 1)) >> 1
}
