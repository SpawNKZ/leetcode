package leetcode

func ReverseString(s []byte) string {
	return string(s[len(s)-1 : 0])
}
