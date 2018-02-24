/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/keyboard-row/description/
 Runtime: 0ms
*/

func findWords(words []string) []string {
	var arr []string

	m := make([]int, 58)

	for _, v := range "QWERTYUIOPqwertyuiop" {
		m[v-'A'] = 1
	}

	for _, v := range "ASDFGHJKLasdfghjkl" {
		m[v-'A'] = 2
	}

	for _, v := range "ZXCVBNMzxcvbnm" {
		m[v-'A'] = 3
	}

	for _, v := range words {
		flag := true
		n := 0

		for _, r := range v {
			if n == 0 {
				n = m[r-'A']
			} else {
				if m[r-'A'] != n {
					flag = false

					break
				}
			}
		}

		if flag {
			arr = append(arr, v)
		}
	}

	return arr
}
