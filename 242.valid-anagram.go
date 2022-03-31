package leetcode

func IsAnagram(s string, t string) bool {
	runes1 := []rune(s)
	runes2 := []rune(t)
	for i := 0; i < len(runes1); i++ {
		for j := i + 1; j < len(runes1); j++ {
			if runes1[i] > runes1[j] {
				runes1[i], runes1[j] = runes1[j], runes1[i]
			}
		}
	}
	for i := 0; i < len(runes2); i++ {
		for j := i + 1; j < len(runes2); j++ {
			if runes2[i] > runes2[j] {
				runes2[i], runes2[j] = runes2[j], runes2[i]
			}
		}
	}
	return string(runes1) == string(runes2)
}
