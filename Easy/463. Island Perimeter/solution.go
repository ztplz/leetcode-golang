/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/island-perimeter/description/
 Runtime: 116ms
*/

func islandPerimeter(grid [][]int) int {
	islandcount := 0
	repeatline := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				islandcount++
			} else {
				continue
			}

			if i-1 >= 0 && grid[i-1][j] == 1 {
				repeatline++
			}

			if i+1 < len(grid) && grid[i+1][j] == 1 {
				repeatline++
			}

			if j-1 >= 0 && grid[i][j-1] == 1 {
				repeatline++
			}

			if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
				repeatline++
			}
		}
	}

	sum := islandcount*4 - repeatline

	return sum
}

