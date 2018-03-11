/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/rotate-string/description/
 Runtime: 0ms
*/

import "strings"

func rotateString(A string, B string) bool {
	db := B + B

	if strings.Contains(db, A) {
		return true
	}

	return false
}