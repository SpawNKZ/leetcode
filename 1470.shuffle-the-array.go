package leetcode

func Shuffle(nums []int, n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, nums[i], nums[i+n])
	}
	return res
}
