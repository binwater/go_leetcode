package main

import (
	"fmt"
	"sort"
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
	value []int
	index int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	tmp := []int{}
	return MinStack{tmp,0}
}


func (this *MinStack) Push(val int)  {
	this.value = append(this.value, val)
	this.index++
}


func (this *MinStack) Pop()  {
	if len(this.value) > 0 {
		this.index--
		this.value = this.value[:this.index]
	}
}


func (this *MinStack) Top() int {
	tmp := 0
	if len(this.value) > 0 {
		tmpIndex := this.index-1
		tmp = this.value[tmpIndex]
	}
	return tmp
}


func (this *MinStack) GetMin() int {
	tmp := make([]int, len(this.value))
	tmpValue := *this
	copy(tmp, tmpValue.value)
	sort.Ints(tmp)
	if len(tmp) > 0 {
		return tmp[0]
	}
	return 0
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
	fmt.Println(minStack.value)
	//minStack.getMin();   --> 返回 -3.
	tmpMin := minStack.GetMin()
	fmt.Println(tmpMin, minStack.value)
	//minStack.pop();
	//minStack.top();      --> 返回 0.
	//minStack.getMin();   --> 返回 -2.
	minStack.Pop()
	fmt.Println("pop is", minStack.value)
	tmpMin = minStack.Top()
	fmt.Println(tmpMin)
	fmt.Println(minStack.value)
	tmpMin = minStack.GetMin()
	fmt.Println(tmpMin)

}