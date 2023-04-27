package main

import "fmt"

/*假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶



提示：

1 <= n <= 45

通过次数1,120,398
提交次数2,073,685

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dpArr := make([]int, n+1)
	for j :=0; j < n+1; j++{
		dpArr[j] = 1
	}

	for i := 2; i <= n; i++{
		dpArr[i] = dpArr[i-1] + dpArr[i-2]
	}
	return dpArr[n]
}

func main() {
	n := 2
	ret := climbStairs(n)
	fmt.Println("ret is: ", ret)

	n = 3
	ret = climbStairs(n)
	fmt.Println("ret is: ", ret)
}
