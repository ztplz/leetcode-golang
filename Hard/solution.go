/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/count-of-smaller-numbers-after-self/description/
 Runtime: 920ms
*/

func countSmaller(nums []int) []int {
	arr := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		count := 0

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}

		arr = append(arr, count)
	}

	return arr
}