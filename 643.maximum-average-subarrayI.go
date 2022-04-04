package leetcode

func FindMaxAverage(nums []int, k int) float64 {
	var sum float64
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	res := sum
	for i := k; i < len(nums); i++ {
		res = res - float64(nums[i-k]) + float64(nums[i])
		if sum < res {
			sum = res
		}
	}
	return sum / float64(k)
}
