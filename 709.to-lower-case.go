package leetcode

func ToLowerCase(s string) string {
	runes := []rune(s)
	var res string
	for _, v := range runes {
		if v >= 65 && v <= 90 {
			v += 32
		}
		res += string(v)
	}
	return res
}
