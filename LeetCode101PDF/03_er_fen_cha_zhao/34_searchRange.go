package main

import (
	"fmt"
	"sort"
)

/*给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。



示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：

输入：nums = [], target = 0
输出：[-1,-1]



提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109*/


func searchRange1(nums []int, target int) []int {
	first, last := 0, 0
	firstFlag, lastFlag := false, false
	for k, v := range nums{
		if v == target{
			first = k
			firstFlag = true
			break
		}
	}

	if !firstFlag{
		first = -1
	}

	i := len(nums)-1
	for ; i>=0; i--{
		if nums[i] == target{
			last = i
			lastFlag = true
			break
		}
	}

	if !lastFlag {
		last = -1
	}
	return []int{first, last}
}

func searchRange(nums []int, target int) []int {
	fmt.Println("nums is: ", nums, " target is: ", target)


	leftmost := sort.SearchInts(nums, target)
	fmt.Println("left most is: ", leftmost)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}

	fmt.Println("tmp value is: ",sort.SearchInts(nums, target + 1) )
	rightmost := sort.SearchInts(nums, target + 1) - 1
	fmt.Println("right most is: ", rightmost)
	fmt.Println()
	return []int{leftmost, rightmost}
}


func main(){
	arr := []int{5,7,7,8,8,10, 11, 12}
	target := 8
	fmt.Println(searchRange(arr, target))

	arr = []int{5,7,7,8,8,10}
	target = 6
	fmt.Println(searchRange(arr, target))

	arr = []int{}
	target = 0
	fmt.Println(searchRange(arr, target))
}

