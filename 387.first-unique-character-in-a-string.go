package leetcode

func FirstUniqChar(s string) int {
	maps := make(map[rune]int)
	runes := []rune(s)
	for _, v := range runes {
		maps[v]++
	}
	for i := 0; i < len(runes); i++ {
		if maps[runes[i]] == 1 {
			return i
		}
	}
	return -1
}
