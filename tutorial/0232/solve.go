package _232

type MyQueue struct {
	st1, st2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		st1: make([]int, 0),
		st2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.st1 = append(this.st1, x)
}

func (this *MyQueue) Pop() int {
	//pop from st1 and push into st2
	tmp := make([]int, 0, len(this.st1))
	for i := len(this.st1) - 1; i >= 0; i-- {
		tmp = append(tmp, this.st1[i])
	}
	this.st1 = []int{}
	this.st2 = append(tmp, this.st2...)
	ans := this.st2[len(this.st2)-1]
	this.st2 = this.st2[:len(this.st2)-1]
	return ans
}

func (this *MyQueue) Peek() int {
	if len(this.st2) > 0 {
		return this.st2[len(this.st2)-1]
	}
	if len(this.st1) > 0 {
		return this.st1[0]
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	return len(this.st1)+len(this.st2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
