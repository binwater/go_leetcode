package main

import (
	"fmt"
	"math"
)

//给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
//返回被除数 dividend 除以除数 divisor 得到的商。
//
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//
//
//
//示例 1:
//
//输入: dividend = 10, divisor = 3
//输出: 3
//解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
//
//示例 2:
//
//输入: dividend = 7, divisor = -3
//输出: -2
//解释: 7/-3 = truncate(-2.33333..) = -2
//
//
//
//提示：
//
//被除数和除数均为 32 位有符号整数。
//除数不为 0。
//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/divide-two-integers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func divide1(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	tmp := float64(dividend)/float64(divisor)
	fmt.Println(tmp)
	result := int(math.Floor(tmp))
	fmt.Println(result)
	if result < 0 {
		if divisor == 1 || divisor == -1 {
			if result > math.MaxInt32 {
				return math.MaxInt32
			}
			return result
		} else {
			if result+1 > math.MaxInt32 {
				return math.MaxInt32
			}
			return result+1
		}
	} else {
		if result > math.MaxInt32 {
			return math.MaxInt32
		}
		return result
	}
}

func divide(dividend int, divisor int) int {
	// 判断被除数与除数的正负性
	positive := true
	if dividend < 0 {
		positive = !positive
		dividend = -dividend
	}
	if divisor < 0 {
		positive = !positive
		divisor = -divisor
	}

	// 计算商
	quotient := 0
	for dividend >= divisor {
		dividend -= divisor
		quotient++
		if quotient > math.MaxInt32 {
			quotient = math.MaxInt32
		}
	}

	if !positive {
		return -quotient
	}
	return quotient
}

func main() {
	dividend := 10
	divisor := 3
	fmt.Printf("divide is: %v\n", divide(dividend, divisor))

	dividend = 1
	divisor = -1
	fmt.Printf("divide is: %v\n", divide(dividend, divisor))
}