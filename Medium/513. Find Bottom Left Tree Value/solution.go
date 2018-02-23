/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/find-bottom-left-tree-value/description/
 Runtime: 12ms
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findBottomLeftValue(root *TreeNode) int {
	var nodearr []*TreeNode

	nodearr = append(nodearr, root)
	ret := 0

	for len(nodearr) != 0 {
		l := len(nodearr)
		ret = nodearr[0].Val
		for i := 0; i < l; i++ {
			if nodearr[i].Left != nil {
				nodearr = append(nodearr, nodearr[i].Left)
			}

			if nodearr[i].Right != nil {
				nodearr = append(nodearr, nodearr[i].Right)
			}
		}

		nodearr = nodearr[l:]
	}

	return ret
}