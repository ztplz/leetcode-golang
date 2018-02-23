/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/maximum-binary-tree/description/
 Runtime: 104ms
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	maxindex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxindex] {
			maxindex = i
		}
	}

	lt := constructMaximumBinaryTree(nums[:maxindex])
	rt := constructMaximumBinaryTree(nums[maxindex+1:])

	newnode := &TreeNode{
		Val:   nums[maxindex],
		Left:  lt,
		Right: rt,
	}

	return newnode
}