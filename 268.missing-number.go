package leetcode

func missingNumber(nums []int) int {
	last := len(nums)
	result := last * (last + 1) / 2
	var SumOfArray int
	for _, v := range nums {
		SumOfArray += v
	}
	return result - SumOfArray
}
