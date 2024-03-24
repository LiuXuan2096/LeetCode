package from_forceRecursive_to_DP

import "math"

/*
给定3个参数，N，M，K
怪兽有N滴血，等着英雄来砍自己
英雄每一次打击，都会让怪兽流失[0~M]的血量
到底流失多少？每一次在[0~M]上等概率的获得一个值
求K次打击之后，英雄把怪兽砍死的概率
*/
func KillMonster(N, M, K int) float64 {
	if N < 1 || M < 1 || K < 1 {
		return 0
	}
	all := int64(math.Pow(float64(M+1), float64(K)))
	kill := process_KillMonster(K, M, N)
	return float64(kill) / float64(all)
}

// 怪兽还剩hp点血
// 每次的伤害在[0~M]范围上
// 还有times次可以砍
// 返回砍死的情况数！
func process_KillMonster(times, M, hp int) int64 {
	if times == 0 {
		if hp <= 0 {
			return 1
		}
		return 0
	}
	if hp <= 0 {
		return int64(math.Pow(float64(M+1), float64(times)))
	}
	var ways int64
	for i := 0; i <= M; i++ {
		ways += process_KillMonster(times-1, M, hp-i)
	}
	return ways
}

func KillMonster_dp(N, M, K int) float64 {
	if N < 1 || M < 1 || K < 1 {
		return 0
	}
	all := int64(math.Pow(float64(M+1), float64(K)))
	dp := make([][]int64, K+1)
	for i := range dp {
		dp[i] = make([]int64, N+1)
	}
	dp[0][0] = 1

	for times := 1; times <= K; times++ {
		dp[times][0] = int64(math.Pow(float64(M+1), float64(times)))
		for hp := 1; hp <= N; hp++ {
			dp[times][hp] = dp[times][hp-1] + dp[times-1][hp]
			if hp-1-M >= 0 {
				dp[times][hp] -= dp[times-1][hp-1-M]
			} else {
				dp[times][hp] -= int64(math.Pow(float64(M+1), float64(times-1)))
			}
		}
	}
	kill := dp[K][N]
	return float64(kill) / float64(all)
}
