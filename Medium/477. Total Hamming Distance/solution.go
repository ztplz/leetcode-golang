/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/total-hamming-distance/description/
 Runtime: 52ms
*/

func totalHammingDistance(nums []int) int {
	l, count := len(nums), 0

	for i := 0; i < 32; i++ {
		bits := 0

		for j := 0; j < l; j++ {
			if nums[j]&1 == 1 {
				bits++
			}

			nums[j] = nums[j] >> 1
		}

		count += bits * (l - bits)
	}

	return count
}

