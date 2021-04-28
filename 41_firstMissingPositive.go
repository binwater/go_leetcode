package main

import (
	"fmt"
	"sort"
)

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
//
//
//进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？
//
//
//
//示例 1：
//
//输入：nums = [1,2,0]
//输出：3
//
//示例 2：
//
//输入：nums = [3,4,-1,1]
//输出：2
//
//示例 3：
//
//输入：nums = [7,8,9,11,12]
//输出：1
//
//
//
//提示：
//
//0 <= nums.length <= 300
//-231 <= nums[i] <= 231 - 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/first-missing-positive
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 1 {
		if nums[0] != 1 {
			return 1
		} else {
			return 2
		}
	}
	first := 0
	i := 0
	for i=0; i<len(nums); i++{
		if nums[i] > 0{
			first = i
			break
		}
	}
	if i == len(nums) {
		return 1
	}

	for i=0; i<len(nums)-1; i++{
		if nums[i] > 0{
			if nums[first] != 1 {
				return 1
			}
			if nums[i+1] - nums[i] != 1 {
				return nums[i]+1
			} else if i+1 == len(nums)-1{
				return nums[i+1]+1
			}
		}
	}
	if i == len(nums)-1 && len(nums) != 2{
		return nums[i]+1
	}else {
		if nums[i] !=1 {
			return 1
		}else {
			return 2
		}
	}

	return 0
}

func main() {
	//nums := []int{1,2,0}
	//fmt.Println(firstMissingPositive(nums))
	//
	//nums = []int{3,4,-1,1}
	//fmt.Println(firstMissingPositive(nums))
	//
	//nums = []int{7,8,9,11,12}
	//fmt.Println(firstMissingPositive(nums))
	//
	//nums = []int{2}
	//fmt.Println(firstMissingPositive(nums))

	nums := []int{2,1}
	fmt.Println(firstMissingPositive(nums))

}