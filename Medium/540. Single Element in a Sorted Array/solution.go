/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/single-element-in-a-sorted-array/description/
 Runtime: 4ms
*/

func singleNonDuplicate(nums []int) int {
	ret := 0

	for _, v := range nums {
		ret = ret ^ v
	}

	return ret
}