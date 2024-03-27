package Go_LeetCode

type MyStack struct {
	queue []int
}

func Constructor_225() MyStack {
	return MyStack{queue: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	i := len(this.queue) - 1 //判断长度
	for i != 0 {
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, val)
		i--
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
