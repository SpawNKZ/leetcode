package leetcode

func DefangIPaddr(address string) string {
	runes := []rune(address)
	var res string
	for _, v := range runes {
		switch v {
		case 46:
			res += "[.]"
		default:
			res += string(v)
		}
	}
	return res
}
