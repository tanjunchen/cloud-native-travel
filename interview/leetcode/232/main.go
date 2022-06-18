package main

/***
"题目：**232 用栈实现队列**

[用栈实现队列](https://leetcode-cn.com/problems/implement-queue-using-stacks/)

请你仅使用两个栈实现先入先出队列。队列应当支持一般队列的支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false

你只能使用标准的栈操作 - 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
*/

type MyQueue struct {
	firstStack  []int
	secondStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.firstStack = append(this.firstStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.secondStack) == 0 {
		for len(this.firstStack) > 0 {
			this.secondStack = append(this.secondStack, this.firstStack[len(this.firstStack)-1])
			this.firstStack = this.firstStack[:len(this.firstStack)-1]
		}
	}
	x := this.secondStack[len(this.secondStack)-1]
	this.secondStack = this.secondStack[:len(this.secondStack)-1]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.secondStack) == 0 {
		for len(this.firstStack) > 0 {
			this.secondStack = append(this.secondStack, this.firstStack[len(this.firstStack)-1])
			this.firstStack = this.firstStack[:len(this.firstStack)-1]
		}
	}
	return this.secondStack[len(this.secondStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.secondStack) == 0 && len(this.firstStack) == 0
}

/**
* Your MyQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Peek();
* param_4 := obj.Empty();
 */

func main() {

}
