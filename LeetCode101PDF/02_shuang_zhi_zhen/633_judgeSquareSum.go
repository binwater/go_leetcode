package main

import (
	"fmt"
	"math"
)

/*给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。



示例 1：

输入：c = 5
输出：true
解释：1 * 1 + 2 * 2 = 5

示例 2：

输入：c = 3
输出：false



提示：

0 <= c <= 231 - 1

通过次数127,828
提交次数332,735

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sum-of-square-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func judgeSquareSum1(c int) bool {
	for a:=0; a*a<=c; a++{
		b := math.Sqrt(float64((c-a*a)))
		if b == math.Floor(b){
			return true
		}
	}
	return false
}

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))

	for left <= right{
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		}else {
			left++
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(3))
}