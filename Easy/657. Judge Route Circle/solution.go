/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/judge-route-circle/description/
 Runtime: 4ms
*/import "strings"

func judgeCircle(moves string) bool {
	return strings.Count(moves, "U") == strings.Count(moves, "D") && strings.Count(moves, "L") == strings.Count(moves, "R")
}