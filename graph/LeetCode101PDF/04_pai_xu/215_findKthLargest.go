package main

import "fmt"

/*给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4



提示：

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/kth-largest-element-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	target := len(nums) - k

	for l <= r {
		div := findDiv1(nums, l, r)
		if div == target{
			return nums[div]
		}
		if div < target{
			l = div + 1
		} else {
			r = div - 1
		}

	}

	return 0
}

func quickSort(arr []int, p, q int){
	//终止条件
	if p > q {
		return
	}

	//查找分区位置
	div := findDiv1(arr, p, q)

	//递归查找
	quickSort(arr, p, div-1)
	quickSort(arr, div+1, q)
}

func findDiv1(arr []int, p, q int) int {
	index := p
	divValue := arr[q]

	for i:=p; i<q; i++{
		if arr[i] < divValue {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[q], arr[index] = arr[index], arr[q]
	return index
}

func findDiv2(arr []int, p, q int) int {
	left, right := make([]int, 0), make([]int, 0)
	divValue := arr[q]
	for i:=p; i<q; i++{
		if arr[i] < divValue {
			left = append(left, arr[i])
		} else  {
			right = append(right, arr[i])
		}
	}

	copy(arr[p:p+len(left)], left)
	arr[p+len(left)] = divValue
	copy(arr[p+len(left)+1:q+1], right)
	return p + len(left)
}

func main() {
	arr := []int{3,2,1,5,6,4}
	k := 2
	ret := findKthLargest(arr, k)
	fmt.Println("ret is: ", ret)

	arr = []int{3,2,3,1,2,4,5,5,6}
	k = 4
	ret = findKthLargest(arr, k)
	fmt.Println("ret is: ", ret)

	arr = []int{1}
	k = 1
	ret = findKthLargest(arr, k)
	fmt.Println("ret is: ", ret)
}

