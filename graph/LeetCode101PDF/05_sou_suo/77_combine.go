package main

import "fmt"

/*给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。



示例 1：

输入：n = 4, k = 2
输出：
[
[2,4],
[3,4],
[2,3],
[1,2],
[1,3],
[1,4],
]

示例 2：

输入：n = 1, k = 1
输出：[[1]]



提示：

1 <= n <= 20
1 <= k <= n

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	comb := make([]int, k)
	count := 0

	backtracking1(&ret, comb, count, 1, n, k)
	return ret
}

func backtracking1(ans *[][]int, comb []int, count, pos, n, k int){
	if count == k {
		tmpArr := make([]int, len(comb))
		copy(tmpArr, comb)
		*ans = append(*ans, tmpArr)
		return
	}

	for i := pos; i<=n; i++{
		comb[count] = i
		count = count + 1
		backtracking1(ans, comb, count, i+1, n, k)
		count = count - 1
	}
}

func main()  {
	n := 4
	k := 2
	ret := combine(n, k)
	fmt.Println("ret is: ", ret)
}

