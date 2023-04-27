package main

import "fmt"

/*给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。



示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
注意，你可以重复使用字典中的单词。

示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false



提示：

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅有小写英文字母组成
wordDict 中的所有字符串 互不相同

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i:=1; i<=n; i++{
		for _,v := range wordDict{
			len:=len(v)
			if i>=len && SubStr(s, i-len, len)==v{
				dp[i] = dp[i] || dp[i-len]
			}
		}
	}
	return dp[n]
}

// 截取字符串，支持多字节字符
// start：起始下标，负数从从尾部开始，最后一个为-1
// length：截取长度，负数表示截取到末尾
func SubStr(str string, start int, length int) (result string) {
	s := []rune(str)
	total := len(s)
	if total == 0 {
		return
	}
	// 允许从尾部开始计算
	if start < 0 {
		start = total + start
		if start < 0 {
			return
		}
	}
	if start > total {
		return
	}
	// 到末尾
	if length < 0 {
		length = total
	}

	end := start + length
	if end > total {
		result = string(s[start:])
	} else {
		result = string(s[start:end])
	}

	return
}


func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	ret := wordBreak(s, wordDict)
	fmt.Println("ret is: ", ret)

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	ret = wordBreak(s, wordDict)
	fmt.Println("ret is: ", ret)

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	ret = wordBreak(s, wordDict)
	fmt.Println("ret is: ", ret)
}