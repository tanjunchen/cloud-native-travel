package main

import "container/list"

/***
"题目：**用两个栈实现队列**

[用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof)

题目描述：
***/

/**
解法一
说明：
**/
type CQueue struct {
	s1, s2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		s1: list.New(),
		s2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s2.Len() == 0 {
		for this.s1.Len() > 0 {
			this.s2.PushBack(this.s1.Remove(this.s1.Back()))
		}
	}
	if this.s2.Len() != 0 {
		res := this.s2.Back()
		this.s2.Remove(res)
		return res.Value.(int)
	}
	return -1
}
