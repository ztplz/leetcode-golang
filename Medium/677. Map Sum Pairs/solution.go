/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/map-sum-pairs/description/
 Runtime: 12ms
*/import "strings"

type MapSum struct {
	Item map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{Item: map[string]int{}}
}

func (this *MapSum) Insert(key string, val int) {
	this.Item[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0

	for k, v := range this.Item {
		if strings.HasPrefix(k, prefix) {
			sum += v
		}
	}

	return sum
}
