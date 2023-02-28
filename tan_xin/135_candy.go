package main

import "fmt"

/*n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。

请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。



示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。

示例 2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。



提示：

n == ratings.length
1 <= n <= 2 * 104
0 <= ratings[i] <= 2 * 104

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/candy
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	arrLength := len(ratings)

	ret := 0
	resultSlice := make([]int, len(ratings))
	for index, _ := range resultSlice {
		resultSlice[index] = 1
	}

	for i := 0; i < arrLength-1; i++ {
		if ratings[i+1] > ratings[i] {
			resultSlice[i+1] = resultSlice[i] + 1
		}
	}

	for j := arrLength - 1; j > 0; j-- {
		if (ratings[j-1] > ratings[j]) && (resultSlice[j-1] <= resultSlice[j]) {
			resultSlice[j-1] = resultSlice[j] + 1
		}
	}

	for _, value := range resultSlice {
		ret = ret + value
	}

	return ret
}

func main() {
/*	fmt.Println("practice 135 ...")

	child := []int{1, 0, 2}
	fmt.Println("first value is: ", candy(child))

	child = []int{1, 2, 2}
	fmt.Println("second value is: ", candy(child))*/

	child := []int{1,3,4,5,2}
	 //1 2 3 5 1
	//expect 11 actual 9
	fmt.Println("third value is: ", candy(child))

}
