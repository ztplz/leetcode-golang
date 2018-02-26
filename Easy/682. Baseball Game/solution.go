/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/baseball-game/description/
 Runtime: 0ms
*/

func calPoints(ops []string) int {
	sum := 0

	roundstack := &slicestack{
		items: make([]int, 0),
	}

	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "C":
			roundstack.pop()
		case "D":
			roundstack.double()
		case "+":
			roundstack.add()
		default:
			num, _ := strconv.ParseInt(ops[i], 10, 64)
			roundstack.push(int(num))
		}
	}

	for j := 0; j < len(roundstack.items); j++ {
		sum += roundstack.items[j]
	}

	return sum
}

type slicestack struct {
	items []int
}

func (ss *slicestack) IsEmpty() bool {
	return len(ss.items) == 0
}

func (ss *slicestack) push(item int) {
	ss.items = append(ss.items, item)
}

func (ss *slicestack) pop() {
	if ss.IsEmpty() {
		return
	}

	ss.items = ss.items[:len(ss.items)-1]
}

func (ss *slicestack) double() {
	if ss.IsEmpty() {
		ss.items = append(ss.items, 0)
	}

	ss.items = append(ss.items, ss.items[len(ss.items)-1]*2)
}

func (ss *slicestack) add() {
	l := len(ss.items)

	if l == 0 {
		ss.items = append(ss.items, 0)
	} else if l == 1 {
		ss.items = append(ss.items, ss.items[0])
	} else {
		ss.items = append(ss.items, ss.items[l-1]+ss.items[l-2])
	}
}
