package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var gLen, sLen = len(g), len(s)
	var count, sIndex = 0, 0
	for i:=0; i< gLen; i++{
		for j:=sIndex; j<sLen; j++{
			if s[j] >= g[i] {
				sIndex = j+1
				count++
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Printf("children\n")
	var g = []int{1,2,3}
	var s = []int{1,1}
	count := findContentChildren(g, s)
	fmt.Printf("output is: %v\n", count)

	g = []int{10,9,8,7}
	s = []int{5,6,7,8}
	count = findContentChildren(g, s)
	fmt.Printf("output is: %v\n", count)
}
