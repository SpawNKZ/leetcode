package leetcode

func isPerfectSquare(num int) bool {
	low := 1
	high := num
	for low <= high {
		med := (low + high) / 2
		if med*med == num {
			return true
		} else if med*med < num {
			low = med + 1
		} else {
			high = med - 1
		}
	}
	return false
}
