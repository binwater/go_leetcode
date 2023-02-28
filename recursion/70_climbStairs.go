package main

import "fmt"

func climbStairs(n int) int {
	if n <=3 {
		return n
	}

	return climbStairs(n-1)+climbStairs(n-2)
}

func main() {
	//输入：n = 2
	//输出：2
	//
	//输入：n = 3
	//输出：3
	//解释：有三种方法可以爬到楼顶。
	//1. 1 阶 + 1 阶 + 1 阶
	//2. 1 阶 + 2 阶
	//3. 2 阶 + 1 阶

	fmt.Println(climbStairs(3))

}
