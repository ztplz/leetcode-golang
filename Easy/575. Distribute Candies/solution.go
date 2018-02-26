/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/distribute-candies/description/
 Runtime: 292ms
*/import "sort"

func distributeCandies(candies []int) int {
	l := len(candies)

	sort.Ints(candies)
	count := 1

	for i := 0; i < l-1; i++ {
		if candies[i] != candies[i+1] {
			count++
		}
	}

	if l/2 <= count {
		return l / 2
	} else {
		return count
	}
}


