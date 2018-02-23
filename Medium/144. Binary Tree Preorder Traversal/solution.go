/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/binary-tree-preorder-traversal/description/
 Runtime: 0ms
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	arr := make([]int, 0)
	nodestack := make([]*TreeNode, 0)
	nodestack = append(nodestack, root)

	for len(nodestack) != 0 {
		cul := nodestack[len(nodestack)-1]
		nodestack = nodestack[:len(nodestack)-1]
		arr = append(arr, cul.Val)

		if cul.Right != nil {
			nodestack = append(nodestack, cul.Right)
		}

		if cul.Left != nil {
			nodestack = append(nodestack, cul.Left)
		}
	}

	return arr
}