package Go_LeetCode

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	nums := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		nums[i] = romanMap[s[i]]
	}

	sum := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			sum -= nums[i]
		} else {
			sum += nums[i]
		}
	}
	return sum + nums[len(nums)-1]
}
