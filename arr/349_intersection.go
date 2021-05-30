package main

import "fmt"

//给定两个数组，编写一个函数来计算它们的交集。
//
//
//
//示例 1：
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//
//示例 2：
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//
//
//
//说明：
//
//输出结果中的每个元素一定是唯一的。
//我们可以不考虑输出结果的顺序。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func intersection(nums1 []int, nums2 []int) []int {
	tmpMap := make(map[int]bool, 0)
	for _, v := range nums1{
		tmpMap[v] = true
	}
	tmpRespNums := make([]int, 0)
	for _, v := range nums2{
		if _, ok := tmpMap[v]; ok{
			tmpRespNums = append(tmpRespNums, v)
			delete(tmpMap, v)
		}
	}
	return tmpRespNums
}

//easy
func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersection(nums1, nums2))

	nums1 = []int{4,9,5}
	nums2 = []int{9,4,9,8,4}
	fmt.Println(intersection(nums1, nums2))
}
