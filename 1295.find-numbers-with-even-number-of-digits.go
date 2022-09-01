package leetcode

func FindNumbers(nums []int) int {
	var res int
	for _, v := range nums {
		if lenOfNum(v)%2 == 0 {
			res++
		}
	}
	return res
}

func lenOfNum(num int) int {
	var counter int
	for num != 0 {
		num /= 10
		counter++
	}
	return counter
}
