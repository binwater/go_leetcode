package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func isValid(s string) bool {
	if len(s) == 0{
		return true
	}

	stackList := []string{}
	judgeMap := map[string]string{")":"(", "}":"{", "]":"["}
	for i:=0; i<len(s); i++{
		if len(stackList) == 0{
			stackList = append(stackList, string(s[i]))
		} else {
			if stackList[len(stackList)-1] == judgeMap[string(s[i])] {
				stackList = stackList[:len(stackList)-1]
			} else {
				stackList = append(stackList, string(s[i]))
			}
		}
	}

	if len(stackList) != 0{
		return false
	} else {
		return true
	}
}

//easy
func main() {
	test := "()"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "()[]{}"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "(]"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "([)]"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "{[]}"
	fmt.Printf("judge is: %v\n", isValid(test))
	test = "{[qdag]}"
	fmt.Printf("judge is: %v\n", isValid(test))
}