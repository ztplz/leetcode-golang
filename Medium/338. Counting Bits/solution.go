/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/counting-bits/description/
 Runtime: 1936ms
*/

func countBits(num int) []int {
	arr := []int{0, 1}

	for i := 2; i <= num; i++ {
		arr = append(arr, arr[i>>1]+i&1)
	}

	return arr[:num+1]
}

