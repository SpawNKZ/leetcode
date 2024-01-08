package leetcode

import (
	"fmt"
)

func RomanToInt(s string) int {
	var res int

	n := map[int]int{73: 1, 86: 5, 88: 10, 76: 50, 67: 100, 68: 500, 77: 1000}

	runes := []rune(s)

	fmt.Println("len(resun", len(runes))
	var last bool

	for i := 0; i < len(runes); i++ {
		if runes[i] == 73 && runes[i+1] == 86 {
			res += 4
			i++
			last = true
		} else if runes[i] == 73 && runes[i+1] == 88 {
			res += 9
			i++
			last = true
		} else if runes[i] == 88 && runes[i+1] == 76 {
			res += 40
			i++
			last = true
		} else if runes[i] == 88 && runes[i+1] == 67 {
			res += 90
			i++
			last = true
		} else if runes[i] == 67 && runes[i+1] == 68 {
			res += 400
			i++
			last = true
		} else if runes[i] == 67 && runes[i+1] == 77 {
			res += 900
			i++
			last = true
		} else {
			res += n[int(runes[i])]
			last = false
		}
		fmt.Println("res", i, res, runes[i])
	}

	fmt.Println("symbols", n, last)

	return res
}
