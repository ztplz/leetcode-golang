/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 Runtime: 8ms
*/

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)

	return nums[len(nums)-k]
}
