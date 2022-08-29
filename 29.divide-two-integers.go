package leetcode

func divide(dividend int, divisor int) int {
	if dividend/divisor >= 1<<31-1 {
		return 1<<31 - 1
	} else if dividend/divisor <= -1<<31 {
		return -1 << 31
	}
	return dividend / divisor
}
