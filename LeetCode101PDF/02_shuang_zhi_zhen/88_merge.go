package main

import (
	"fmt"
	"sort"
)

/*给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。



示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。

示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。

示例 3：

输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。



提示：

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func merge1(nums1 []int, m int, nums2 []int, n int)  {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

func merge3(nums1 []int, m int, nums2 []int, n int)  {
	for p1, p2, tail := m-1, n-1, m+n-1; p1>=0 || p2>=0; tail--{
		var cur int
		if p1==-1{
			cur = nums2[p2]
			p2--
		} else if p2==-1{
			cur = nums1[p1]
			p1--
		} else if nums1[p1]>nums2[p2]{
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

func main(){
	//nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	//合并结果是 [1,2,2,3,5,6]
	a := []int{1,2,3,0,0,0}
	merge(a, 3, []int{2,5,6}, 3)
	fmt.Println(a)


	//nums1 = [1], m = 1, nums2 = [], n = 0
	//合并结果是 [1]
	a = []int{1}
	merge(a, 1, []int{}, 0)
	fmt.Println(a)

	//nums1 = [0], m = 0, nums2 = [1], n = 1
	//合并结果是 [1]
	a = []int{0}
	merge(a, 0, []int{1}, 1)
	fmt.Println(a)
}