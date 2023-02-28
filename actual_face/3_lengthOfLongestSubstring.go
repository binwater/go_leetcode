package actual_face

import "fmt"

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//示例 4:
//
//输入: s = ""
//输出: 0
//
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s 由英文字母、数字、符号和空格组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func lengthOfLongestSubstring1(s string) int {
	if s == ""{
		return 0
	}
	var max, index int
	uS := []rune(s)
	for k, v := range uS{
		for i:=index;i<k;i++{
			if v==uS[i]{
				index = i+1
			}
		}

		if k-index+1 > max{
			max = k-index+1
		}
	}

	return max
}

func lengthOfLongestSubstring(s string) int {
	var max, index = 0, 0
	bStr := []byte(s)
	for k, v := range bStr{
		for i:=index; i<k; i++{
			if bStr[i] == v{
				index = i+1
			}
		}

		if k - index +1 > max {
			max = k - index +1
		}
	}
	return max
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))

	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))

	s = ""
	fmt.Println(lengthOfLongestSubstring(s))
}
