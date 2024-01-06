package Go_LeetCode

func SsumSubarrayMins(arr []int) int {
	var stack = make([]int, len(arr))
	var left = nearLessEqualLeft(arr, stack)
	var right = nearLessRight(arr, stack)
	var ans int64
	for i := 0; i < len(arr); i++ {
		var start = i - left[i]
		var end = right[i] - i
		ans += int64(start * end * arr[i])
		ans %= mod
	}
	return int(ans)
}

func nearLessEqualLeft(arr []int, stack []int) []int {
	var N = len(arr)
	var left = make([]int, N)
	var size = 0
	for i := N - 1; i >= 0; i-- {
		for size != 0 && arr[i] <= arr[stack[size-1]] {
			size--
			left[stack[size]] = i
		}
		stack[size] = i
		size++
	}
	for size != 0 {
		size--
		left[stack[size]] = -1
	}
	return left
}

func nearLessRight(arr []int, stack []int) []int {
	var N = len(arr)
	var right = make([]int, N)
	var size = 0
	for i := 0; i < N; i++ {
		for size != 0 && arr[i] < arr[stack[size-1]] {
			size--
			right[stack[size]] = i
		}
		stack[size] = i
		size++
	}
	for size != 0 {
		size--
		right[stack[size]] = N
	}
	return right
}

const mod = 1000000007

func GPTsumSubarrayMins(arr []int) int {
	stack := make([]int, len(arr))
	left := GPTnearLessEqualLeft(arr, stack)
	right := GPTnearLessRight(arr, stack)

	var ans int64
	for i := 0; i < len(arr); i++ {
		start := int64(i - left[i])
		end := int64(right[i] - i)
		ans += start * end * int64(arr[i])
		ans %= mod
	}
	return int(ans)
}

func GPTnearLessEqualLeft(arr []int, stack []int) []int {
	N := len(arr)
	left := make([]int, N)
	size := 0

	for i := N - 1; i >= 0; i-- {
		for size != 0 && arr[i] <= arr[stack[size-1]] {
			left[stack[size-1]] = i
			size--
		}
		stack[size] = i
		size++
	}

	for size != 0 {
		left[stack[size-1]] = -1
		size--
	}

	return left
}

func GPTnearLessRight(arr []int, stack []int) []int {
	N := len(arr)
	right := make([]int, N)
	size := 0

	for i := 0; i < N; i++ {
		for size != 0 && arr[stack[size-1]] > arr[i] {
			right[stack[size-1]] = i
			size--
		}
		stack[size] = i
		size++
	}

	for size != 0 {
		right[stack[size-1]] = N
		size--
	}

	return right
}
