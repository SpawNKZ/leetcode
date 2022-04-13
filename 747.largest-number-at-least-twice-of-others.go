package leetcode

func DominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var max int
	var index int
	for i, v := range nums {
		if max < v {
			max, index = v, i
		}
	}
	for _, v := range nums {
		if v != max && v*2 > max {
			return -1
		}
	}
	return index
}
