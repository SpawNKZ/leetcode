package leetcode

func BuildArray(nums []int) []int {
	var arr []int
	for i := 0; i < len(nums); i++ {
		arr = append(arr, nums[nums[i]])
	}
	return arr
}
