package graph

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	var row [9][10]byte
	var col [9][10]byte
	var box [9][10]byte

	for i:=0; i<9; i++{
		for j:=0; j<9; j++{
			if board[i][j] == '.'{
				continue
			}
			curNmu := board[i][j]-'0'
			if row[i][curNmu]==1{
				return false
			}
			if col[j][curNmu]==1{
				return false
			}
			if box[j/3 +(i/3)*3][curNmu]==1{
				return false
			}

			row[i][curNmu]=1
			col[j][curNmu]=1
			box[j/3 +(i/3)*3][curNmu]=1
		}
	}
	return true
}

func main() {
	board :=[][]byte{{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'}}

	fmt.Println(isValidSudoku(board))
}