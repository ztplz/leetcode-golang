/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/jewels-and-stones/description/
 Runtime: 0ms
*/

func numJewelsInStones(J string, S string) int {
	var count int
	m := make(map[rune]int)

	for _, v1 := range J {
		m[v1]++
	}

	for _, v2 := range S {
		if _, ok := m[v2]; ok {
			count++
		}
	}

	return count
}