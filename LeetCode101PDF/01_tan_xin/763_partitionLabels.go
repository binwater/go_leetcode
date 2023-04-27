package main

import "fmt"

/*给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。

注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。

返回一个表示每个字符串片段的长度的列表。


示例 1：

输入：s = "ababcbaca defegde hijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。

示例 2：

输入：s = "eccbbbbdec"
输出：[10]



提示：

1 <= s.length <= 500
s 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-labels
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/


func partitionLabels(s string) (partition []int) {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	fmt.Println(lastPos)

	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}


func main() {
	s := "ababcbacadefegdehijhklij"

	partitionArr := partitionLabels(s)
	fmt.Println("input is: ", s)
	fmt.Println("ret is: ", partitionArr)
	for i, v := range partitionArr{
		fmt.Println("index is: ", i, " value is: ", v)
	}
}