package Go_LeetCode

// O(N)的解
// 因为是正数，所以不用考虑符号位(31位)
// 首先来到30位，假设剩余的数字有N个(整体)，看看这一位是1的数，有几个
// 如果有0个、或者1个
// 说明不管怎么在数组中选择，任何两个数&的结果在第30位上都不可能有1了
// 答案在第30位上的状态一定是0，
// 保留剩余的N个数，继续考察第29位，谁也不淘汰(因为谁也不行，干脆接受30位上没有1的事实)
// 如果有2个，
// 说明答案就是这两个数(直接返回答案)，因为别的数在第30位都没有1，就这两个数有。
// 如果有>2个，比如K个
// 说明答案一定只用在这K个数中去选择某两个数，因为别的数在第30位都没有1，就这K个数有。
// 答案在第30位上的状态一定是1，
// 只把这K个数作为剩余的数，继续考察第29位，其他数都淘汰掉
// .....
// 现在来到i位，假设剩余的数字有M个，看看这一位是1的数，有几个
// 如果有0个、或者1个
// 说明不管怎么在M个数中选择，任何两个数&的结果在第i位上都不可能有1了
// 答案在第i位上的状态一定是0，
// 保留剩余的M个数，继续考察第i-1位
// 如果有2个，
// 说明答案就是这两个数(直接返回答案)，因为别的数在第i位都没有1，就这两个数有。
// 如果有>2个，比如K个
// 说明答案一定只用在这K个数中去选择某两个数，因为别的数在第i位都没有1，就这K个数有。
// 答案在第i位上的状态一定是1，
// 只把这K个数作为剩余的数，继续考察第i-1位，其他数都淘汰掉

func MaxAndValue(arr []int) int {
	M := len(arr)
	ans := 0
	for bit := 30; bit >= 0; bit-- {
		i := 0
		tmp := M
		for i < M {
			if (arr[i] & (1 << bit)) == 0 {
				swap_07_01(arr, i, M-1)
				M--
			} else {
				i++
			}
		}
		if M == 2 {
			return arr[0] & arr[1]
		}
		if M < 2 {
			M = tmp
		} else {
			ans |= (1 << bit)
		}
	}
	return ans
}

func swap_07_01(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
