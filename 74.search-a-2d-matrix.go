package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			switch {
			case target > matrix[i][j]:
				continue
			case target < matrix[i][j]:
				return false
			default:
				return true
			}
		}
	}
	return false
}
