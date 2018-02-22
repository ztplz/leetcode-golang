/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/reverse-string/description/
 Runtime: 8ms
*/import "strings"

func reverseString(s string) string {
	strArr := strings.Split(s, "")
	l := len(strArr)

	for i := 0; i < l/2; i++ {
		strArr[i], strArr[l-i-1] = strArr[l-i-1], strArr[i]
	}

	str := strings.Join(strArr, "")

	return str
}