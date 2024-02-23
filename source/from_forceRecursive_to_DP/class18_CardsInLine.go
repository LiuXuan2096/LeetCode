package from_forceRecursive_to_DP

/*
题目描述：
给定一个整型数组arr，代表数值不同的纸牌排成一条线
玩家A和玩家B依次拿走每张纸牌
规定玩家A先拿，玩家B后拿，谁最后分数大谁赢
但是每个玩家每次只能拿走最左或最右的纸牌
玩家A和玩家B都绝顶聪明，请返回最后获胜者的分数
*/

// 最原始的暴力递归版本的解法
func Win1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	first := f1(arr, 0, len(arr)-1)
	second := g1(arr, 0, len(arr)-1)
	return max(first, second)
}

// 	// arr[L..R]，先手获得的最好分数返回
func f1(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	p1 := arr[L] + g1(arr, L+1, R)
	p2 := arr[R] + g1(arr, L, R-1)
	return max(p1, p2)
}

// // arr[L..R]，后手获得的最好分数返回
func g1(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	p1 := f1(arr, L+1, R)
	p2 := f1(arr, L, R-1)
	return min(p1, p2)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 从暴力递归走向动态规划的第一步：傻缓存法
func Win2(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	N := len(arr)
	fmap := make([][]int, N) // 先手函数f2的缓存
	gmap := make([][]int, N) // 后手函数g2的缓存
	for i := 0; i < N; i++ {
		fmap[i] = make([]int, N)
		gmap[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmap[i][j] = -1
			gmap[i][j] = -1
		}
	}
	first := f2(arr, 0, len(arr)-1, fmap, gmap)
	second := g2(arr, 0, len(arr)-1, fmap, gmap)

	return max(first, second)
}

// arr[L..R]，先手获得的最好分数返回
func f2(arr []int, L, R int, fmap, gmap [][]int) int {
	if fmap[L][R] != -1 {
		return fmap[L][R]
	}
	var ans int
	if L == R {
		ans = arr[L]
	} else {
		p1 := arr[L] + g2(arr, L+1, R, fmap, gmap)
		p2 := arr[R] + g2(arr, L, R-1, fmap, gmap)
		ans = max(p1, p2)
	}
	fmap[L][R] = ans
	return ans
}

// arr[L..R]，后手获得的最好分数返回
func g2(arr []int, L, R int, fmap, gmap [][]int) int {
	if gmap[L][R] != -1 {
		return gmap[L][R]
	}
	var ans int
	if L != R {
		p1 := f2(arr, L+1, R, fmap, gmap)
		p2 := f2(arr, L, R-1, fmap, gmap)
		ans = min(p1, p2)
	}
	gmap[L][R] = ans
	return ans
}

// Win3 动态规划的最终版本
func Win3(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	N := len(arr)
	fmap := make([][]int, N)
	gmap := make([][]int, N)
	for i := 0; i < N; i++ {
		fmap[i] = make([]int, N)
		gmap[i] = make([]int, N)
		fmap[i][i] = arr[i]
	}

	// 从对角线开始依次填充 两张缓存表
	for startCol := 1; startCol < N; startCol++ {
		L, R := 0, startCol
		for R < N {
			fmap[L][R] = max(arr[L]+gmap[L+1][R], arr[R]+gmap[L][R-1])
			gmap[L][R] = min(fmap[L+1][R], fmap[L][R-1])
			L++
			R++
		}
	}
	return max(fmap[0][N-1], gmap[0][N-1])
}
