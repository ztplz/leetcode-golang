/*
 Author:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/hamming-distance/description/
 Runtime: 0ms
*/

func hammingDistance(x int, y int) int {
	n := x ^ y
	count := 0

	for n != 0 {
		if n&1 == 1 {
			count++
		}

		n = n >> 1
	}

	return count
}