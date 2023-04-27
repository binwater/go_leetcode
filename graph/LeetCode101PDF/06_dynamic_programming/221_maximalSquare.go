package main

import (
	"fmt"
	"math"
	"sort"
)

/*在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。



示例 1：

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4

示例 2：

输入：matrix = [["0","1"],["1","0"]]
输出：1

示例 3：

输入：matrix = [["0"]]
输出：0



提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximal-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for k:=0; k<m+1; k++{
		dp[k] = make([]int, n+1)
	}

	maxSide := 0
	for i := 1; i<=m; i++{
		for j := 1; j<=n; j++{
			if matrix[i-1][j-1] == '1'{
				tmpArr := []int{dp[i-1][j-1], dp[i][j-1], dp[i-1][j]}
				sort.Ints(tmpArr)
				dp[i][j] = tmpArr[0]+1
			}
			maxSide= int(math.Max(float64(maxSide), float64(dp[i][j])))
		}
	}

	return maxSide*maxSide
}

func main() {
	matrix := [][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	ret := maximalSquare(matrix)
	fmt.Println("ret is: ", ret)
}
