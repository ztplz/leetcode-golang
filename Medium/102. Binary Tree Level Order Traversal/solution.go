/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 Runtime: 4ms
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	nodearr := make([]*TreeNode, 0)
	arr := make([][]int, 0)
	nodearr = append(nodearr, root)

	for len(nodearr) != 0 {
		l := len(nodearr)
		subarr := make([]int, l)
		for i := 0; i < l; i++ {
			subarr[i] = nodearr[i].Val

			if nodearr[i].Left != nil {
				nodearr = append(nodearr, nodearr[i].Left)
			}

			if nodearr[i].Right != nil {
				nodearr = append(nodearr, nodearr[i].Right)
			}
		}

		nodearr = nodearr[l:]
		arr = append(arr, subarr)
	}

	return arr
}

