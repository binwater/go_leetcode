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



提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

通过次数387,068
提交次数558,870

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func minPathSum(grid [][]int) int {
	for i:=0; i<len(grid); i++{
		for j:=0; j<len(grid[0]); j++{
			//当左边和上边都是矩阵边界时： 即当i=0,j=0i = 0, j = 0i=0,j=0时，其实就是起点， dp[i][j]=grid[i][j]dp[i][j] = grid[i][j]dp[i][j]=grid[i][j]；
			if i==0 && j==0{
				continue
			} else if i==0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j==0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else{
				grid[i][j] = int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1]))) + grid[i][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Println(minPathSum(grid))
}
