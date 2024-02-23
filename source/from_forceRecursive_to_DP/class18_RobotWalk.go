package from_forceRecursive_to_DP

/*
假设有排成一行的N个位置记为1~N，N一定大于或等于2
开始时机器人在其中的M位置上(M一定是1~N中的一个)
如果机器人来到1位置，那么下一步只能往右来到2位置；
如果机器人来到N位置，那么下一步只能往左来到N-1位置；
如果机器人来到中间位置，那么下一步可以往左走或者往右走；
规定机器人必须走K步，最终能来到P位置(P也是1~N中的一个)的方法有多少种
给定四个参数 N、M、K、P，返回方法数
*/

// Ways1 最原始 最容易想到的暴力递归的解法
func Ways1(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}
	return process1(start, K, aim, N)
}

// process1 当前机器人来到的当前位置是cur
// 机器人还有rest步要走，最终的目标是aim
// 有哪些位置？ 1-N
// 返回值：机器人从cur出发，走过rest步之后，最终停在aim的方法数，是多少？
func process1(cur, rest, aim, N int) int {
	if rest == 0 {
		if cur == aim {
			return 1
		}
		return 0
	}

	if cur == 1 {
		return process1(2, rest-1, aim, N)
	}

	if cur == N {
		return process1(N-1, rest-1, aim, N)
	}

	return process1(cur-1, rest-1, aim, N) + process1(cur+1, rest-1, aim, N)
}

// Ways2 从暴力递归向动态规划进阶的第一步：加入傻缓存
func Ways2(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}

	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	// dp就是缓存表
	// dp[cur][rest] == -1 -> process1(cur, rest)之前没算过！
	// dp[cur][rest] != -1 -> process1(cur, rest)之前算过！返回值，dp[cur][rest]
	// N+1 * K+1
	return process2(start, K, aim, N, dp)
}

// process2
// cur 范: 1 ~ N
// rest 范：0 ~ K
func process2(cur, rest, aim, N int, dp [][]int) int {
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}

	// 执行到这里，说明当前位置以前没算过
	var ans int
	if rest == 0 {
		if cur == aim {
			ans = 1
		} else {
			ans = 0
		}
	} else if cur == 1 {
		ans = process2(2, rest-1, aim, N, dp)
	} else if cur == N {
		ans = process2(N-1, rest-1, aim, N, dp)
	} else {
		ans = process2(cur-1, rest-1, aim, N, dp) + process2(cur+1, rest-1, aim, N, dp)
	}

	dp[cur][rest] = ans
	return ans

}

// Ways3 动态规划的版本
func Ways3(N, start, aim, K int) int {
	if N < 2 || start < 1 || start > N || aim < 1 || aim > N || K < 1 {
		return -1
	}

	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}
	dp[aim][0] = 1
	for rest := 1; rest <= K; rest++ {
		dp[1][rest] = dp[2][rest-1]
		dp[N][rest] = dp[N-1][rest-1]
		for cur := 2; cur < N; cur++ {
			dp[cur][rest] = dp[cur-1][rest-1] + dp[cur+1][rest-1]
		}
	}

	return dp[start][K]
}
