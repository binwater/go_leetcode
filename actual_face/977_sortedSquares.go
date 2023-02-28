package actual_face

import (
	"fmt"
	"math"
	"sort"
)

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//
//示例 2：
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
//
//
//
//提示：
//
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums 已按 非递减顺序 排序
//
//
//
//进阶：
//
//请你设计时间复杂度为 O(n) 的算法解决本问题
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func sortedSquares1(nums []int) []int {
   tmpNums := make([]int, len(nums))
   for i:=0; i<len(nums); i++ {
   	tmpNums[i] = nums[i] * nums[i]
   }

   sort.Ints(tmpNums)
   return tmpNums
}

func sortedSquares(nums []int) []int {
	tmpNums := make([]int, len(nums))
	var left = 0
	var right = len(nums) - 1
	var idx = len(nums) - 1

	for left <= right {
		if math.Abs(float64(nums[left])) > float64(nums[right]) {
			tmpNums[idx] = nums[left] * nums[left]
			left++
		} else {
			tmpNums[idx] = nums[right] * nums[right]
			right--
		}
		idx--
	}
	return tmpNums
}

func main() {
	nums := []int{-4,-1,0,3,10}
	res := sortedSquares(nums)
	fmt.Printf("res is: %v\n", res)
}