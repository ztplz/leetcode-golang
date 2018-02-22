/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/array-partition-i/description/
 Runtime: 80ms
*/

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	n := len(nums) / 2
	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[2*i]
	}

	return sum
}
