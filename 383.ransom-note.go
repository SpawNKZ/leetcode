package leetcode

func CanConstruct(ransomNote string, magazine string) bool {
	runes1 := []rune(ransomNote)
	runes2 := []rune(magazine)
	maps1 := make(map[rune]int)
	maps2 := make(map[rune]int)
	for _, v := range runes1 {
		maps1[v]++
	}
	for _, v := range runes2 {
		maps2[v]++
	}
	keys := make([]rune, 0, len(maps1))
	for k := range maps1 {
		keys = append(keys, k)
	}
	var res int
	for _, v := range keys {
		if maps1[v] <= maps2[v] {
			res++
		}
	}
	return res == len(maps1)
}
