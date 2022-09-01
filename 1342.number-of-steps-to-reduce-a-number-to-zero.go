package leetcode

func NumberOfSteps(num int) int {
	var res int
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		res++
	}
	return res
}
