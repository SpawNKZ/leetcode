package leetcode

func CanConstruct(ransomNote string, magazine string) bool {
	runes1 := []rune(ransomNote)
	runes2 := []rune(magazine)
	var res int
	var counter int
	for i := 0; i < len(runes1); i++ {
		for j := counter; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				res++
				counter++
				break
			}
			counter++
		}
	}
	return res == len(runes1)
}
