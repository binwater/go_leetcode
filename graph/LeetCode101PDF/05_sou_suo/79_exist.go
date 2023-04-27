package main

import "fmt"

/*给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

示例 2：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true

示例 3：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false



提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}

	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	find := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++{
			backtrackingExist(i, j, board, word, &find, visited, 0)
		}
	}

	return find
}

func backtrackingExist(i, j int, board [][]byte, word string, find *bool, visited [][]bool, pos int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}

	if visited[i][j] || *find || board[i][j] != word[pos] {
		return
	}

	if pos == len(word)-1 {
		*find = true
		return
	}

	visited[i][j] = true

	backtrackingExist(i+1, j, board, word, find, visited, pos+1)
	backtrackingExist(i-1, j, board, word, find, visited, pos+1)
	backtrackingExist(i, j+1, board, word, find, visited, pos+1)
	backtrackingExist(i, j-1, board, word, find, visited, pos+1)

	visited[i][j] = false
}

func main() {
/*	输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
	输出：true*/

	board := [][]byte{{[]byte("A")[0],[]byte("B")[0],[]byte("C")[0],[]byte("E")[0]},
		{[]byte("S")[0],[]byte("F")[0],[]byte("C")[0],[]byte("S")[0]},
		{[]byte("A")[0],[]byte("D")[0],[]byte("E")[0],[]byte("E")[0]}}
	word := "ABCCED"

	ret := exist(board, word)
	fmt.Println("ret is: ", ret)
}
