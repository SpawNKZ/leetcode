package leetcode

func FindMiddleIndex(nums []int) int {
	var index int
	var counter int
	for counter != len(nums) {
		var sum1 int
		var sum2 int
		for i := 0; i < index; i++ {
			sum1 += nums[i]
		}
		for j := index + 1; j < len(nums); j++ {
			sum2 += nums[j]
		}
		if sum1 == sum2 {
			return index
		}
		index++
		counter++
	}
	return -1
}
