/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
package leetcode

func MaxSubArray(nums []int) int {
	max := nums[0]
	var sum int
	for _, v := range nums {
		sum += v
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

// @lc code=end
