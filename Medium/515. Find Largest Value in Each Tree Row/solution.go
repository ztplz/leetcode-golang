/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/
 Runtime: 16ms
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func largestValues(root *TreeNode) []int {
	var nodearr []*TreeNode
	var arr []int

	if root == nil {
		return nil
	}

	nodearr = append(nodearr, root)

	for len(nodearr) != 0 {
		l := len(nodearr)
		max := nodearr[0].Val

		for i := 0; i < l; i++ {
			if nodearr[i].Val >= max {
				max = nodearr[i].Val
			}

			if nodearr[i].Left != nil {
				nodearr = append(nodearr, nodearr[i].Left)
			}

			if nodearr[i].Right != nil {
				nodearr = append(nodearr, nodearr[i].Right)
			}
		}

		nodearr = nodearr[l:]
		arr = append(arr, max)
	}

	return arr
}