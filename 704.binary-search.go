package leetcode

// Search Binary
func Search(nums []int, target int) int {
	var right, left, mid int

	left = 0
	right = len(nums) - 1

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right -= 1
		} else {
			left += 1
		}
	}

	return -1
}
