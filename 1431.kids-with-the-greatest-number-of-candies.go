package leetcode

func KidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, 0, len(candies))
	var maxNumberInCandies int

	for i := 0; i < len(candies); i++ {
		if candies[i] > maxNumberInCandies {
			maxNumberInCandies = candies[i]
		}
	}
	for _, v := range candies {
		if v+extraCandies >= maxNumberInCandies {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}

	return res
}
