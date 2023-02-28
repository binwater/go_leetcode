package greed_div_recall_dp

import "fmt"

//老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
//
//那么这样下来，老师至少需要准备多少颗糖果呢？
//
//
//
//示例 1：
//
//输入：[1,0,2]
//输出：5
//解释：你可以分别给这三个孩子分发 2、1、2 颗糖果。
//
//示例 2：
//
//输入：[1,2,2]
//输出：4
//解释：你可以分别给这三个孩子分发 1、2、1 颗糖果。
//第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/candy
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func candy(ratings []int) int {
	if ratings == nil || len(ratings) == 0 {
		return 0
	}

	n := len(ratings)
	reward := make([]int,n)

	// 初始化每个人所分糖果，并比较左值
	reward[0] = 1
	res := 1
	for i:=1;i<n;i++{
		if ratings[i]>ratings[i-1]{
			reward[i] = reward[i-1]+1
		}else{
			reward[i] = 1
		}
		res += reward[i]
	}

	// 比较右值，使得分最高者糖果最多
	for i:=n-2;i>=0;i--{
		diff := reward[i+1]+1-reward[i]
		if ratings[i]>ratings[i+1] && diff> 0{
			reward[i] += diff
			res += diff
		}
	}

	return res
}

func main() {
	nums := []int{1, 0, 2}
	count := candy(nums)
	fmt.Printf("count is: %v\n", count)

	nums = []int{1, 2, 2}
	count = candy(nums)
	fmt.Printf("count is: %v\n", count)

	nums = []int{1,3,2,2,1}
	count = candy(nums)
	fmt.Printf("count is: %v\n", count)
}