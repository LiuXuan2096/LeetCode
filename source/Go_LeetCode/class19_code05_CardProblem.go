package Go_LeetCode

/*
 * 一张扑克有3个属性，每种属性有3种值（A、B、C）
 * 比如"AAA"，第一个属性值A，第二个属性值A，第三个属性值A
 * 比如"BCA"，第一个属性值B，第二个属性值C，第三个属性值A
 * 给定一个字符串类型的数组cards[]，每一个字符串代表一张扑克
 * 从中挑选三张扑克，一个属性达标的条件是：这个属性在三张扑克中全一样，或全不一样
 * 挑选的三张扑克达标的要求是：每种属性都满足上面的条件
 * 比如："ABC"、"CBC"、"BBC"
 * 第一张第一个属性为"A"、第二张第一个属性为"C"、第三张第一个属性为"B"，全不一样
 * 第一张第二个属性为"B"、第二张第二个属性为"B"、第三张第二个属性为"B"，全一样
 * 第一张第三个属性为"C"、第二张第三个属性为"C"、第三张第三个属性为"C"，全一样
 * 每种属性都满足在三张扑克中全一样，或全不一样，所以这三张扑克达标
 * 返回在cards[]中任意挑选三张扑克，达标的方法数
 *
 * */

/*
func main() {
	cards := []string{"ABA", "BAA", "AAC", "ABA", "ABA", "AAC", "AAC"}
	fmt.Println(Go_LeetCode.Ways_19_05(cards)) // 输出: true
}
*/

func Ways_19_05(cards []string) int {
	counts := make([]int, 27)
	for _, s := range cards {
		str := []rune(s)
		// 统计每种类型的牌的数量
		counts[int(str[0]-'A')*9+(int(str[1]-'A')*3)+(int(str[2]-'A'))*1]++
	}
	ways := 0
	// 在相同类型的牌中选择达标的方法数
	for status := 0; status < 27; status++ {
		n := counts[status]
		if n > 2 {
			if n == 3 {
				ways += 1
			} else {
				ways += n * (n - 1) * (n - 2) / 6
			}
		}
	}
	path := make([]int, 0)
	for i := 0; i < 27; i++ {
		if counts[i] != 0 {
			path = append(path, i)
			ways += process_19_05(counts, i, &path)
			path = path[:len(path)-1]
		}
	}
	return ways
}

func process_19_05(counts []int, pre int, path *[]int) int {
	if len(*path) == 3 {
		return getWays(counts, *path)
	}
	ways := 0
	for next := pre + 1; next < 27; next++ {
		if counts[next] != 0 {
			*path = append(*path, next)
			ways += process_19_05(counts, next, path)
			*path = (*path)[:len(*path)-1]
		}
	}
	return ways
}

// getWays 判断当前选择的牌是否达标
func getWays(counts []int, path []int) int {
	v1 := path[0]
	v2 := path[1]
	v3 := path[2]
	for i := 9; i > 0; i /= 3 {
		cur1 := v1 / i
		cur2 := v2 / i
		cur3 := v3 / i
		v1 %= i
		v2 %= i
		v3 %= i
		if (cur1 != cur2 && cur1 != cur3 && cur2 != cur3) || (cur1 == cur2 && cur2 == cur3) {
			continue
		}
		return 0
	}
	v1 = path[0]
	v2 = path[1]
	v3 = path[2]
	return counts[v1] * counts[v2] * counts[v3]
}
