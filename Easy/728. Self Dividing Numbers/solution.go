/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/self-dividing-numbers/description/
 Runtime: 0ms
*/

func selfDividingNumbers(left int, right int) []int {
	var numArr []int

	for i := left; i <= right; i++ {
		if isSelfNum(i) {
			numArr = append(numArr, i)
		}
	}

	return numArr
}

func isSelfNum(a int) bool {
	b := a

	for a != 0 {
		c := a % 10
		if c == 0 {
			return false
		}

		if b%c != 0 {
			return false
		}

		a /= 10
	}

	return true
}