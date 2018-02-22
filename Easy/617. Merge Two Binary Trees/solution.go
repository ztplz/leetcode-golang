/*
 Author:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/merge-two-binary-trees/description/
 Runtime: 48ms
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	newnode := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	if t1 != nil && t2 != nil {
		newnode.Val = t1.Val + t2.Val
		newnode.Left = mergeTrees(t1.Left, t2.Left)
		newnode.Right = mergeTrees(t1.Right, t2.Right)
	} else if t1 != nil && t2 == nil {
		newnode = t1
	} else {
		newnode = t2
	}

	return newnode
}