package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	sideLength := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sideLength = dfs(grid, i, j)
				break
			}
		}
	}
	return sideLength
}
func dfs(grid [][]int, i, j int) int {
	if !(i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])) {
		return 1
	}

	if grid[i][j]==0{
		return 1
	}

	if grid[i][j] != 1{
		return 0
	}

	grid[i][j] = 2
	return dfs(grid, i+1, j) + dfs(grid, i-1, j)+ dfs(grid, i, j+1)+ dfs(grid, i, j-1)
}

func main() {
	//输入：grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
	//输出：16
	//解释：它的周长是上面图片中的 16 个黄色的边

	input := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(input))
}
