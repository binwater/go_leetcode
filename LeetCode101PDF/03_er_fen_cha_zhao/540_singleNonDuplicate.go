package main

import "fmt"

/*给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。



示例 1:

输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2

示例 2:

输入: nums =  [3,3,7,7,10,11,11]
输出: 10



提示:

1 <= nums.length <= 105
0 <= nums[i] <= 105

通过次数112,219
提交次数185,173

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/single-element-in-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums) -1
	for low < high {
		mid := (low+high)/2
		fmt.Println("mid index is: ", mid ^ 1)
		if nums[mid] == nums[mid ^ 1]{
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func main()  {
	nums := []int{1,1,2,3,3,4,4,8,8}
	fmt.Println(singleNonDuplicate(nums))

	nums = []int{3,3,7,7,10,11,11}
	fmt.Println(singleNonDuplicate(nums))
}