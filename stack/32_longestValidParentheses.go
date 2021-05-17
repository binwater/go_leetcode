package main

import "fmt"

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
// 
//
//示例 1：
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//
//输入：s = ""
//输出：0
// 
//
//提示：
//
//0 <= s.length <= 3 * 104
//s[i] 为 '(' 或 ')'
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func longestValidParentheses(s string) int {
	maxlenth := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0;i < len(s);i ++ {
		if s[i] == '(' {
			stack = append(stack, i)
		}else {
			// 先出栈 然后再判断是否为空，空的话说明没有"("，所以就把")"入栈，如果没有空，那么就和前面的索引相减得到长度。
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			}else {
				maxlenth = max(maxlenth,i - stack[len(stack)-1])
			}
		}
	}
	return maxlenth
}

func max(a int, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

//hard
func main() {
	s := "(()"
	fmt.Println(longestValidParentheses(s))

	s = ")()())"
	fmt.Println(longestValidParentheses(s))

	s = ""
	fmt.Println(longestValidParentheses(s))
}