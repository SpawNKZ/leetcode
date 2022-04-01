package leetcode

func NumJewelsInStones(jewels string, stones string) int {
	runes1 := []rune(jewels)
	runes2 := []rune(stones)
	var res int
	for _, v := range runes1 {
		for _, val := range runes2 {
			if v == val {
				res++
			}
		}
	}
	return res
}
