/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/trim-a-binary-search-tree/description/
 Runtime: 24ms
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	newnode := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	if root.Val > R {
		newnode = trimBST(root.Left, L, R)

		return newnode
	} else if root.Val < L {
		newnode = trimBST(root.Right, L, R)

		return newnode
	} else {
		newnode.Val = root.Val
		newnode.Left = trimBST(root.Left, L, R)
		newnode.Right = trimBST(root.Right, L, R)
	}

	return newnode
}