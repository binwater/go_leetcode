package main

import (
	"fmt"
	"math"
)

/*给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。



示例 1：

输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
输出：[[0,0,0],[0,1,0],[0,0,0]]

示例 2：

输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
输出：[[0,0,0],[0,1,0],[1,2,1]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/01-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 {
		return nil
	}
	n, m := len(mat), len(mat[0])

	dp := make([][]int, n)
	for k:=0; k<n; k++{
		dp[k] = make([]int, m)
		for index, _ := range dp[k]{
			dp[k][index]=math.MaxInt32
		}
	}

	//从左上到右下
	for i:=0; i<n; i++{
		for j:=0; j<m; j++{
			if mat[i][j]==0{
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i][j-1] + 1)))
				}
				if i > 0 {
					dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i-1][j] + 1)))
				}
			}
		}
	}

	//从右下到左上
	for i := n - 1; i >= 0; i--{
		for j := m - 1; j >= 0; j--{
			if mat[i][j] != 0{
				if j < m -1 {
					dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i][j+1] + 1)))
				}
				if i < n-1{
					dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i+1][j] + 1)))
				}

			}
		}
	}

	return dp
}

func main() {
	mat := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	ret := updateMatrix(mat)
	fmt.Println("ret is: ", ret)

	mat = [][]int{{0,0,0},{0,1,0},{1,1,1}}
	ret = updateMatrix(mat)
	fmt.Println("ret is: ", ret)

	mat = [][]int{{1,1,0,0,1,0,0,1,1,0},
		{1,0,0,1,0,1,1,1,1,1},
		{1,1,1,0,0,1,1,1,1,0},
		{0,1,1,1,0,1,1,1,1,1},
		{0,0,1,1,1,1,1,1,1,0},
		{1,1,1,1,1,1,0,1,1,1},
		{0,1,1,1,1,1,1,0,0,1},
		{1,1,1,1,1,0,0,1,1,1},
		{0,1,0,1,1,0,1,1,1,1},
		{1,1,1,0,1,0,1,1,1,1}}
	ret = updateMatrix(mat)
	fmt.Println("ret is: ", ret)
}