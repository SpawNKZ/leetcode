package leetcode

func IsPalindrome(s string) bool {
	runes := []rune(s)
	var res string
	for _, v := range runes {
		if v >= 65 && v <= 90 {
			res += string(v + 32)
		} else if v >= 97 && v <= 122 {
			res += string(v)
		} else if v >= 48 && v <= 57 {
			res += string(v)
		}
	}
	var b int
	left, right := 0, len(res)-1
	for left < right {
		if res[left] == res[right] {
			b++
			left++
			right--
			continue
		}
		return false
	}
	if len(res)%2 == 0 {
		return b*2 == len(res)
	}
	return b*2+1 == len(res)
}
