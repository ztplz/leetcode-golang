/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/number-complement/description/
 Runtime: 0ms
*/

func findComplement(num int) int {
	s := strconv.FormatInt(int64(num), 2)
	mask := 2

	for i := 0; i < len(s)-1; i++ {
		mask = mask * 2
	}

	return ^num & (mask - 1)
}
