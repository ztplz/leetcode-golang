/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/fizz-buzz/description/
 Runtime: 140ms
*/

func fizzBuzz(n int) []string {
	strArr := make([]string, n)

	for i := 0; i < n; i++ {
		if (i+1)%15 == 0 {
			strArr[i] = "FizzBuzz"
		} else if (i+1)%3 == 0 {
			strArr[i] = "Fizz"
		} else if (i+1)%5 == 0 {
			strArr[i] = "Buzz"
		} else {
			strArr[i] = strconv.Itoa(i + 1)
		}
	}

	return strArr
}
