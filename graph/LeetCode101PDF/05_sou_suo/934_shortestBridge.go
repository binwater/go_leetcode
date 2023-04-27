package main

import "fmt"

/*给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。

岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。

你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。

返回必须翻转的 0 的最小数目。



示例 1：

输入：grid = [[0,1],[1,0]]
输出：1

示例 2：

输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
输出：2

示例 3：

输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
输出：1



提示：

n == grid.length == grid[i].length
2 <= n <= 100
grid[i][j] 为 0 或 1
grid 中恰有两个岛

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/shortest-bridge
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func shortestBridge(grid [][]int) int {
	type pair struct {
		x, y int
	}

	step := 0
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	for i, row := range grid{
		for j, v := range row{
			if v != 1{
				continue
			}
			q := []pair{}
			var dfs func(int, int)
			dfs = func(i, j int) {
				grid[i][j] = -1
				q = append(q, pair{i, j})
				for _, d := range dirs{
					x, y := i+d.x, j+d.y
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						dfs(x, y)
					}
				}
			}
			dfs(i, j)

			for {
				tmp := q
				q = nil
				for _, p := range tmp{
					for _, d := range dirs{
						x, y := p.x+d.x, p.y+d.y
						if 0 <= x && x < n && 0 <= y && y < n {
							if grid[x][y] == 1 {
								return step
							}

							if grid[x][y] == 0{
								grid[x][y] = -1
								q = append(q, pair{x, y})
							}
						}
					}
				}
				step++
			}
		}
	}

	return step
}


func main() {
	arr := [][]int {{0,1,0},{0,0,0},{0,0,1}}
	fmt.Println(shortestBridge(arr))
}
