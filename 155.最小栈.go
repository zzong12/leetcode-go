/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */
// package main

// @lc code=start
type MinStack struct {
	q  []int
	mq []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.q = append(this.q, val)
	if len(this.mq) == 0 {
		this.mq = append(this.mq, val)
	} else {
		min := this.mq[len(this.mq)-1]
		if min > val {
			min = val
		}
		this.mq = append(this.mq, min)
	}
}

func (this *MinStack) Pop() {
	this.q = this.q[0 : len(this.q)-1]
	this.mq = this.mq[0 : len(this.mq)-1]
}

func (this *MinStack) Top() int {
	return this.q[len(this.q)-1]
}

func (this *MinStack) GetMin() int {
	return this.mq[len(this.mq)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
