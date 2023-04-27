package main

import "fmt"

/*给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：

输入：nums = [1]
输出：[[1]]



提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	backtracking(nums, 0, &ret)
	return ret
}

func backtracking(nums []int, level int, ans *[][]int){
	if level == len(nums)-1{
		tmpArr := make([]int, len(nums))
		copy(tmpArr, nums)
		*ans = append(*ans, tmpArr)
		return
	}

	for i := level; i<len(nums); i++{
		nums[i], nums[level] = nums[level], nums[i]
		backtracking(nums, level+1, ans)
		nums[i], nums[level] = nums[level], nums[i]
	}
}

func main(){
	nums := []int{1,2,3}
	ret := permute(nums)
	fmt.Println("ret is: ", ret)
}