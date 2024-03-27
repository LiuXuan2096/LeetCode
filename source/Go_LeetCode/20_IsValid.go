package Go_LeetCode

type Stack struct {
	data []rune
}

func (s *Stack) push(ch rune) {
	s.data = append(s.data, ch)
}

func (s *Stack) pop() rune {
	if len(s.data) == 0 {
		return 0 // 0表示空栈
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func isValid_20(s string) bool {
	if s == "" {
		return true
	}
	str := []rune(s)
	stack := &Stack{}
	for _, cha := range str {
		if cha == '(' || cha == '[' || cha == '{' {
			if cha == '(' {
				stack.push(')')
			} else if cha == '[' {
				stack.push(']')
			} else {
				stack.push('}')
			}
		} else {
			if stack.isEmpty() {
				return false
			}
			last := stack.pop()
			if cha != last {
				return false
			}
		}
	}
	return stack.isEmpty()
}
