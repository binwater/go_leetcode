package main

import (
	"fmt"
)

/*你一个字符串 s，最多 可以从中删除一个字符。

请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。



示例 1：

输入：s = "aba"
输出：true

示例 2：

输入：s = "abca"
输出：true
解释：你可以删除字符 'c' 。

示例 3：

输入：s = "abc"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-palindrome-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func check(s string) bool{
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r]{
			l++
			r--
		}else{
			flag1, flag2 := true, true

			for i, j:=l, r-1; i<j; i, j = i+1, j-1{
				if s[i] != s[j]{
					flag1 = false
					break
				}
			}

			for i, j:=l+1, r; i<j; i, j = i+1, j-1{
				if s[i] != s[j]{
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

func validPalindrome(s string) bool {
	if check(s){
		return true
	}
	return false
}


func main() {
	//s := "aba"
	//fmt.Println(validPalindrome(s))
	//s = "abca"
	//fmt.Println(validPalindrome(s))
	s := "deeee"
	fmt.Println(validPalindrome(s))
}