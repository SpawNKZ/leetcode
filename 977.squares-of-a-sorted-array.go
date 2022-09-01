package leetcode

func SortedSquares(nums []int) []int {
	var arr []int
	for _, v := range nums {
		arr = append(arr, v*v)
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
