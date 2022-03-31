package leetcode

/*brute-force*/
/*func Intersect(nums1 []int, nums2 []int) []int {
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
	for i := 0; i < len(nums2); i++ {
		for j := i + 1; j < len(nums2); j++ {
			if nums2[i] > nums2[j] {
				nums2[i], nums2[j] = nums2[j], nums2[i]
			}
		}
	}
	var arr []int
	var counter int
	// var c int
	for i := 0; i < len(nums1); i++ {
		// counter = 0
		for j := counter; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				counter = j
				counter++
				arr = append(arr, nums1[i])
				// fmt.Println(counter)
				// fmt.Println(counter, arr)
				break
			}
			// fmt.Println(j)

		}
	}
	return arr
}
*/
func Intersect(nums1 []int, nums2 []int) []int {
	maps := make(map[int]int)

	for _, v := range nums1 {
		maps[v] = maps[v] + 1
	}

	var res []int

	for _, v := range nums2 {
		if maps[v] > 0 {
			maps[v] = maps[v] - 1
			res = append(res, v)
		}
	}
	return res
}
