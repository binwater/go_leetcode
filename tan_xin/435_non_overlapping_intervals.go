package main

import (
	"fmt"
	"sort"
)

/*给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。



示例 1:

输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。

示例 2:

输入: intervals = [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。

示例 3:

输入: intervals = [ [1,2], [2,3] ]
输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。



提示:

1 <= intervals.length <= 105
intervals[i].length == 2
-5 * 104 <= starti < endi <= 5 * 104

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/non-overlapping-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	total, prev := 0, intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] < prev {
			total++
		} else {
			prev = intervals[i][1]
		}
	}

	return total
}

func main() {
	//[[1,2],[2,3],[3,4],[1,3]]
	arr := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println("value is: ", eraseOverlapIntervals(arr))

}
