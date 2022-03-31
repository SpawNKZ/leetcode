package leetcode

import "fmt"

func CheckInclusion(s1 string, s2 string) bool {
	runes1 := []rune(s1)
	fmt.Println(runes1)
	runes := []rune(s2)
	fmt.Println(runes)
	var counter int
	var res int
	for i := len(runes1) - 1; i >= 0; i-- {
		fmt.Println("Commin")
		fmt.Println(i)
		fmt.Println(runes1[i])
		for j := counter; j < len(runes); j++ {
			if counter == 10 {
				break
			}
			if runes[j] == runes1[i] {
				counter = j
				res++
				fmt.Println(runes1[i])
				fmt.Println("counter:", counter, "res", res)
				break
			}
			if res >= 1 && runes[j] != runes1[i] {
				fmt.Println("YES")
				res = 0
				i = len(runes1) - 1
				// break
			}
		}
	}
	return res == len(runes1)
}
