/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/next-greater-element-i/description/
 Runtime: 16ms
*/

func nextGreaterElement(findNums []int, nums []int) []int {
	var a []int
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		m[v] = i
	}

	for j := 0; j < len(findNums); j++ {
		flag := true
		findex := m[findNums[j]]

		if findex == len(nums)-1 {
			a = append(a, -1)

			continue
		}

		for k := findex + 1; k < len(nums); k++ {
			if nums[k] > findNums[j] {
				a = append(a, nums[k])
				flag = false

				break
			}
		}

		if flag {
			a = append(a, -1)
		}

	}

	return a
}