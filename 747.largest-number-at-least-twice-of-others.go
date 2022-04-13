package leetcode

func DominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var max int
	var index int
	for i, v := range nums {
		if max < v {
			max = v
			index = i
		}
	}
	for _, v := range nums {
		if v == max {
			continue
		}
		var res int
		res += v * 2
		if max < res {
			return -1
		}
	}
	return index
}
