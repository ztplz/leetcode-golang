/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/reshape-the-matrix/description/
 Runtime: 84ms
*/

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}

	numarr := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			numarr = append(numarr, nums[i][j])
		}
	}

	arr := make([][]int, 0)

	for i := 0; i < r; i++ {
		subarr := make([]int, 0)
		for j := 0; j < c; j++ {
			subarr = append(subarr, numarr[i*c+j])
		}

		arr = append(arr, subarr)
	}

	return arr
}