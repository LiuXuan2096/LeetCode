package Go_LeetCode

/*
给定一个栈，请逆序这个栈，不能申请额外的数据结构，只能使用递归函数
*/

func Reverse(stack *[]int) {
	if len(*stack) == 0 {
		return
	}
	var i = f(stack)
	Reverse(stack)
	*stack = append(*stack, i)
}

// 栈顶元素移出掉
// 上面的元素盖下来
// 返回移除掉的栈顶元素
func f(stack *[]int) int {
	var result = (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	if len(*stack) == 0 {
		return result
	}
	var last = f(stack)
	*stack = append(*stack, result)
	return last
}
