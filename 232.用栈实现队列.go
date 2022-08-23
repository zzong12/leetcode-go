/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */
// package main

// @lc code=start
type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	this.in2out()
	head := this.out[len(this.out)-1]
	this.out = this.out[0 : len(this.out)-1]
	return head
}

func (this *MyQueue) Peek() int {
	this.in2out()
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

func (this *MyQueue) in2out() {
	if len(this.out) == 0 && len(this.in) > 0 {
		for i := len(this.in) - 1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
		}
		this.in = make([]int, 0)
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
