package Go_LeetCode

import (
	"math"
	"strings"
)

func MinStickers(stickers []string, target string) int {
	var N = len(stickers)
	var counts [][26]int = make([][26]int, N)
	for i := 0; i < N; i++ {
		str := []byte(stickers[i])
		for _, ch := range str {
			counts[i][ch-'a']++
		}
	}
	var dp map[string]int = make(map[string]int, 10)
	dp[""] = 0
	var ans = process3(counts, target, &dp)
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}

func process3(stickers [][26]int, t string, dp *map[string]int) int {
	v, ok := (*dp)[t]
	if ok {
		return v
	}
	target := []byte(t)
	var tcounts [26]int
	for _, ch := range target {
		tcounts[ch-'a']++
	}
	var N = len(stickers)
	var min = math.MaxInt32
	for i := 0; i < N; i++ {
		var sticker = stickers[i]
		if sticker[target[0]-'a'] > 0 {
			var builder strings.Builder
			for j := 0; j < 26; j++ {
				if tcounts[j] > 0 {
					nums := tcounts[j] - sticker[j]
					for k := 0; k < nums; k++ {
						builder.WriteByte(byte(j + 'a'))
					}
				}
			}
			rest := builder.String()
			min = int(math.Min(float64(min), float64(process3(stickers, rest, dp))))
		}
	}
	var ans int
	if min == math.MaxInt32 {
		ans = min
	} else {
		ans = min + 1
	}
	(*dp)[t] = ans
	return ans
}

//func process(stickers []string, target string) int {
//	if len(target) == 0 {
//		return 0
//	}
//	var min = math.MaxInt32
//	for _, first := range stickers {
//		var rest = minus(target, first)
//		if len(rest) != len(target) {
//			min = int(math.Min(float64(min), float64(process(stickers, rest))))
//		}
//	}
//	if min == math.MaxInt32 {
//		return min
//	} else {
//		return min + 1
//	}
//
//}
//
//func minus(target, first string) string {
//	str1 := []byte(target)
//	str2 := []byte(first)
//	var count [26]int
//	for _, ch := range str1 {
//		count[ch-'a']++
//	}
//	for _, ch := range str2 {
//		count[ch-'a']--
//	}
//	var b strings.Builder
//	for i := 0; i < 26; i++ {
//		if count[i] > 0 {
//			for j := 0; j < count[i]; j++ {
//				b.WriteByte(byte(i + 'a'))
//			}
//		}
//	}
//	return b.String()
//}
