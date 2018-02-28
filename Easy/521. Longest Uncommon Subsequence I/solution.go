/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/longest-uncommon-subsequence-i/description/
 Runtime: 0ms
*/

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	if len(a) > len(b) {
		return len(a)
	} else {
		return len(b)
	}
}
