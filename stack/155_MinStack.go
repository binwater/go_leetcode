package main

import (
	"fmt"
)

//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) —— 将元素 x 推入栈中。
//pop() —— 删除栈顶的元素。
//top() —— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。
//
//
//示例:
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//提示：
//
//pop、top 和 getMin 操作总是在 非空栈 上调用。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/min-stack
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MinStack struct {
	allStack []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.allStack = append(this.allStack, val)
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if this.minStack[len(this.minStack)-1] == this.allStack[len(this.allStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.allStack = this.allStack[:len(this.allStack)-1]
}

func (this *MinStack) Top() int {
	return this.allStack[len(this.allStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	var minStack = Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.allStack)
	//minStack.getMin();   --> 返回 -3.
	tmpMin := minStack.GetMin()
	fmt.Println(tmpMin, minStack.minStack)
	//minStack.pop();
	minStack.Pop()
	fmt.Println("pop is", minStack.allStack)
	//minStack.top();      --> 返回 0.
	tmpMin = minStack.Top()
	fmt.Println(tmpMin)
	fmt.Println(minStack.allStack)
	//minStack.getMin();   --> 返回 -2.
	tmpMin = minStack.GetMin()
	fmt.Println(minStack.minStack)
	fmt.Println(tmpMin)

}
