package main

import (
	"fmt"
	"leetcode"
)

func main() {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums2 := []int{2, 5, 6}
	// n := 3
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	leetcode.Merge(nums1, m, nums2, n)
	fmt.Println(nums1)
	// fmt.Println(leetcode.MaxSubArray(num1))
	// fmt.Println(leetcode.MaxSubArray(num2))
}
