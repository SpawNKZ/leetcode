package leetcode

/*
Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1,1,1,2,2,3,3,3], k = 2
Output: [1,3]

Input: nums = [1], k = 1
Output: [1]
*/

func TopKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}

	freq := make([][]int, len(nums)+1)

	for k, v := range numMap {
		freq[v] = append(freq[v], k)
	}

	var res []int
	for i := len(freq) - 1; i >= 0; i-- {
		for _, v := range freq[i] {
			if len(res) == k {
				return res
			}
			res = append(res, v)
		}
	}

	return res
}
