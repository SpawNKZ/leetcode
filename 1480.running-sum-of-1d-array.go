package leetcode

func RunningSum(nums []int) []int {
	var res []int
	switch len(nums) {
	case 1:
		return nums
	}
	var a int
	for i := 1; i < len(nums); i++ {
		sum := nums[0]
		for j := 1; j < i; j++ {
			sum += nums[j]
		}
		res = append(res, sum)
		a = sum
	}
	a += nums[len(nums)-1]
	res = append(res, a)
	return res
}
