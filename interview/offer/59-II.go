package main

/***
"题目：**队列的最大值**

[队列的最大值](https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof)

题目描述：请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空, pop_front 和 max_value 需要返回 -1
***/

/**
解法一
说明：
**/
type MaxQueue struct {
	data []int
	max  []int
}

func Constructor2() MaxQueue {
	return MaxQueue{data: make([]int, 0), max: make([]int, 0)}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
	length := len(this.max) - 1
	for length >= 0 {
		if this.max[length] > value {
			break
		}
		length--
	}
	this.max = append(this.max[:length+1], value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}
	res := this.data[0]
	if res == this.max[0] {
		this.max = this.max[1:]
	}
	this.data = this.data[1:]
	return res
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
