package leetcode

func targetIndices(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	var nums1 []int
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			nums1 = append(nums1, i)
		}
	}
	if len(nums1) == 0 {
		return nums1
	} else {
		return nums1
	}
}
