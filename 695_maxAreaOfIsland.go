package main

import "fmt"

/*
给你一个大小为 m x n 的二进制矩阵 grid 。
岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
岛屿的面积是岛上值为 1 的单元格的数目。
计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

示例 1：

输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
输出：6
解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。

示例 2：

输入：grid = [[0,0,0,0,0,0,0,0]]
输出：0

提示：
    m == grid.length
    n == grid[i].length
    1 <= m, n <= 50
    grid[i][j] 为 0 或 1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxAreaOfIsland(grid [][]int) int {
	maxNum := 0
	for i:=0; i<len(grid); i++{
		for j:=0; j<len(grid[0]); j++{
			if grid[i][j]==1{
				ret := getArea(grid, i, j)
				if ret > maxNum{
					maxNum = ret
				}
			}
		}
	}

	return maxNum
}

func getArea(grid [][]int, i, j int) int {
	if i == len(grid)||i <0{
		return 0
	}else if j==len(grid[0])||j<0{
		return 0
	} else {
		if grid[i][j]==1{
			grid[i][j]=0
			return 1+getArea(grid, i+1, j)+getArea(grid, i-1, j)+getArea(grid, i, j+1)+getArea(grid, i, j-1)
		}
	}
	return 0
}

func main(){
	grid := [][]int{{0,0,1,0,0,0,0,1,0,0,0,0,0},{0,0,0,0,0,0,0,1,1,1,0,0,0},{0,1,1,0,1,0,0,0,0,0,0,0,0},{0,1,0,0,1,1,0,0,1,0,1,0,0},{0,1,0,0,1,1,0,0,1,1,1,0,0},{0,0,0,0,0,0,0,0,0,0,1,0,0},{0,0,0,0,0,0,0,1,1,1,0,0,0},{0,0,0,0,0,0,0,1,1,0,0,0,0}}
	fmt.Println(maxAreaOfIsland(grid))
}
