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
	// nums1 := []int{1, 2, 2, 1}
	// nums2 := []int{2, 2}
	// nums1 := []int{1, 2, 2, 1}
	// nums2 := []int{2}
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}
	// fmt.Println(leetcode.Intersect(nums1, nums2))
	s := "anagram"
	t := "nagaram"
	fmt.Println(leetcode.IsAnagram(s, t))
	// fmt.Println(leetcode.MaxSubArray(num1))
	// fmt.Println(leetcode.MaxSubArray(num2))
}
