/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
package leetcode

import "fmt"

func MaxSubArray(nums []int) int {
	var arr []int
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums); i++ {
		sum := i
		for j := i + 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				sum += nums[j]
			} else {
				fmt.Println(sum)
				arr = append(arr, sum)
				break
			}
		}
	}
	Max := arr[0]
	for i := 0; i < len(arr); i++ {
		if Max < arr[i] {
			Max = arr[i]
		}
	}
	return Max
}

// @lc code=end
