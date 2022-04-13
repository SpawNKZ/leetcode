package leetcode

import "fmt"

func PlusOne(digits []int) []int {
	var sum int
	for _, v := range digits {
		sum = sum*10 + v
		fmt.Println(sum)
	}
	sum++
	fmt.Println(sum)
	var res []int
	for sum != 0 {
		res = append([]int{sum % 10}, res...)
		sum /= 10
	}
	return res
}
