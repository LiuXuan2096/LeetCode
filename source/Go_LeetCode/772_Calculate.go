package Go_LeetCode

import "strconv"

func calculate(s string) int {
	return f_772([]rune(s), 0)[0]
}

func f_772(str []rune, i int) []int {
	queue := make([]string, 0)
	cur := 0
	var bra []int
	// 从i出发，开始撸串
	for i < len(str) && str[i] != ')' {
		if str[i] >= '0' && str[i] <= '9' {
			cur = cur*10 + int(str[i]-'0')
			i++
		} else if str[i] != '(' {
			// 遇到的是运算符号
			addNum(&queue, cur, str[i])
			cur = 0
			i++
		} else {
			// 遇到左括号了
			bra = f_772(str, i+1)
			cur = bra[0]
			i = bra[1] + 1
		}
	}
	addNum(&queue, cur, '+')
	return []int{getAns(queue), i}
}

func addNum(queue *[]string, num int, op rune) {
	stack := *queue
	N := len(stack)
	if N > 0 && (stack[N-1] == "*" || stack[N-1] == "/") {
		top := stack[N-1]
		pre, _ := strconv.Atoi(stack[N-2])
		*queue = (*queue)[:len(*queue)-2]
		if top == "*" {
			num *= pre
		} else {
			num = pre / num
		}
	}
	*queue = append(*queue, strconv.Itoa(num), string(op))
}

func getAns(queue []string) int {
	ans, _ := strconv.Atoi(queue[0])
	for i := 1; i < len(queue)-1; i += 2 {
		num, _ := strconv.Atoi(queue[i+1])
		if queue[i] == "+" {
			ans += num
		} else {
			ans -= num
		}
	}
	return ans
}
