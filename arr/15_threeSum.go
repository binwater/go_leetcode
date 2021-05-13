package arr

import (
	"fmt"
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
// 
//
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//示例 2：
//
//输入：nums = []
//输出：[]
//示例 3：
//
//输入：nums = [0]
//输出：[]
// 
//
//提示：
//
//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/3sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		z := len(nums) - 1
		for z > j {
			b := nums[j]
			c := nums[z]

			if nums[i]+b+c > 0 {
				z--
			} else if nums[i]+b+c < 0 {
				j++
			} else {
				item := []int{nums[i], b, c}
				result = append(result, item)

				for j < z && nums[j] == nums[j+1] {
					j++
				}
				for j < z && nums[z] == nums[z-1] {
					z--
				}

				j++
				z--
			}
		}
	}

	return result
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i:=0; i<len(nums)-1; i++{
		if i>0 && nums[i] == nums[i-1]{
			continue
		}

		j := i+1
		z := len(nums)-1
		for j < z{
			b := nums[j]
			c := nums[z]
			if nums[i] + b +c > 0 {
				z--
			} else if nums[i] + b +c < 0 {
				j++
			} else {
				tmp := []int{nums[i], b, c}
				result = append(result, tmp)

				for j < z && nums[j] == nums[j+1]{
					j++
				}
				for j < z && nums[z] == nums[z-1]{
					z--
				}
				j++
				z--
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))

	nums = []int{}
	fmt.Println(threeSum(nums))

	nums = []int{0}
	fmt.Println(threeSum(nums))

	nums = []int{1,-1,-1,0}
	fmt.Println(threeSum(nums))

}