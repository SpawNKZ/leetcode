package leetcode

import "strconv"

func FizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		var str string
		if i%3 == 0 {
			str = "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = strconv.Itoa(i)
		}
		res = append(res, str)
	}
	return res
}
