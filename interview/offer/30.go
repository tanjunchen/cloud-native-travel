package main

/***
"题目：**包含min函数的栈**

[包含min函数的栈](https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof)

题目描述：
***/

/**
解法一
说明：
**/
type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{min: []int{}, stack: []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.min) == 0 || this.min[len(this.min)-1] >= x {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	l := len(this.stack)
	if l > 0 {
		if k := len(this.min); k > 0 && this.min[k-1] == this.stack[l-1] {
			this.min = this.min[:k-1]
		}
		this.stack = this.stack[:l-1]
	}
}

func (this *MinStack) Top() int {
	if l := len(this.stack); l > 0 {
		return this.stack[l-1]
	}
	return 0
}

func (this *MinStack) Min() int {
	if l := len(this.min); l > 0 {
		return this.min[l-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
