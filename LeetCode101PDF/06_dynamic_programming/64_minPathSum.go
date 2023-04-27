package main

import (
	"fmt"
	"math"
)

/*给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。



示例 1：

输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//dp[i][j]表示从左上角开始到（i,j）位置的最优路径的数字和
//因每次只能向下或者向右移动，状态转移方程：
//dp[i][j]=min(dp[i-1][j], dp[i][j-1]) + grid[i][j] 其中grid表示原数组
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for k := 0; k < m; k++{
		dp[k] = make([]int, n)
	}

	for i := 0; i < m; i++{
		for j:=0; j < n; j++{
			if i==0 && j==0{
				dp[i][j] = grid[i][j]
			} else if i==0{
				dp[i][j] = dp[i][j-1]+grid[i][j]
			} else if j==0{
				dp[i][j] = dp[i-1][j]+grid[i][j]
			} else {
				dp[i][j] = int(math.Min(float64( dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	ret := minPathSum(grid)
	fmt.Println("ret is: ", ret)

	grid = [][]int{{1,2,3},{4,5,6}}
	ret = minPathSum(grid)
	fmt.Println("ret is: ", ret)
}