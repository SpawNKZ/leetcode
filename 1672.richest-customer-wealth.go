package leetcode

func MaximumWealth(accounts [][]int) int {
	var res int
	for i := 0; i < len(accounts); i++ {
		var sum int
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if res < sum {
			res = sum
		}
	}
	return res
}
