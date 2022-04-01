package leetcode

func MatrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c < len(mat)*len(mat[0]) {
		return mat
	}
	if r >= len(mat) && c >= len(mat[0]) {
		return mat
	}
	res := [][]int{}
	var arr []int

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			arr = append(arr, mat[i][j])
			if len(arr) >= c {
				res = append(res, arr)
				arr = []int{}
			}
		}
	}
	if len(arr) != 0 {
		res = append(res, arr)
	}
	return res
}
