package main

import (
	"fmt"
)

/*给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。



示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。

示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-equal-subset-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum := 0
	maxNum := 0
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	if maxNum > target {
		return false
	}
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		num := nums[i]
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[target]
}


func main() {
	nums := []int{1,5,11,5}
	ret := canPartition(nums)
	fmt.Println("ret is: ", ret)

	nums = []int{1,2,3,5}
	ret = canPartition(nums)
	fmt.Println("ret is: ", ret)
}
