package leetcode

func FindMaxConsecutiveOnes(nums []int) int {
	var res int
	var a int
	for _, v := range nums {
		if v == 1 {
			res++
		} else {
			if a < res {
				a = res
			}
			res = 0
		}
	}
	if a < res {
		return res
	}
	return a
}
