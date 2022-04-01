package leetcode

func Generate(numRows int) [][]int {
	switch numRows {
	case 1:
		return [][]int{{1}}
	case 2:
		return [][]int{{1}, {1, 1}}
	case 3:
		return [][]int{{1}, {1, 1}, {1, 2, 1}}
	}
	res := [][]int{{1}, {1, 1}, {1, 2, 1}}
	for i := 3; i < numRows; i++ {
		var arr []int
		for k := 0; k < len(res[i-1]); k++ {
			switch k {
			case 0:
				arr = append(arr, 1)
				continue
			}
			for l := k - 1; l < len(res[i-1]); l++ {
				var sum int
				sum = sum + res[i-1][l] + res[i-1][k]
				arr = append(arr, sum)
				//fmt.Println(k)
				//fmt.Println(res[i-1][k])
				break
			}
		}
		arr = append(arr, 1)
		res = append(res, arr)
	}
	return res
}
