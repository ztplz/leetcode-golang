/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/single-number-iii/description/
 Runtime: 12ms
*/

func singleNumber(nums []int) []int {
	var arr []int
	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = 1
		}
	}

	for k, _ := range m {
		arr = append(arr, k)
	}

	return arr
}

