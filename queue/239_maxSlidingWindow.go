package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if nums==nil || k > len(nums){
		return nil
	}
	var ret []int

	for i:=0; i<len(nums); i++{
		fmt.Println("i is: ", i)
		fmt.Println("i+k is: ", i+k)
		ret = append(ret, maxNums(nums[i:i+k]))
		if (i+k)==len(nums){
			break
		}
	}
	return ret
}

func maxNums(arr []int) int{
    max := arr[0]
	for i:=0; i<len(arr); i++{
		if arr[i] > max{
			max = arr[i]
		}
	}
	return max
}

func main(){
	//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
	//输出：[3,3,5,5,6,7]

	//nums := []int{1,3,-1,-3,5,3,6,7}
	//k := 3
	//fmt.Println(maxSlidingWindow(nums, k))

	nums := []int{7, 2, 4}
	k := 2
	fmt.Println(maxSlidingWindow(nums, k))
}