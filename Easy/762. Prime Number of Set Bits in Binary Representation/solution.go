/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/description/
 Runtime: 72ms
*/

func countPrimeSetBits(L int, R int) int {
	var arr []int

	for i := L; i <= R; i++ {
		a := i
		count := 0
		flag := true

		for a > 0 {
			if a&1 == 1 {
				count++
			}

			a = a >> 1
		}

		if count == 0 {
			continue
		}

		if count == 1 {
			continue
		}

		if count == 2 {
			arr = append(arr, count)

			continue
		}

		for j := 2; j < count; j++ {
			if count%j == 0 {
				flag = false

				break
			}
		}

		if flag {
			arr = append(arr, count)
		}
	}

	return len(arr)
}