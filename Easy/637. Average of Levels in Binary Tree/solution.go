/*
 Author: ztplz
 leetcode profile: https://leetcode.com/ztplz/
 Email:     mysticzt@gmail.com
 Url:     https://leetcode.com/problems/average-of-levels-in-binary-tree/description/
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

func averageOfLevels(root *TreeNode) []float64 {
	var treearr [](*TreeNode)
	var arr []float64

	if root == nil {
		return nil
	}

	treearr = append(treearr, root)

	for len(treearr) != 0 {
		l := len(treearr)
		sum := 0.0
		for i := 0; i < l; i++ {
			sum += float64(treearr[0].Val)

			if treearr[0].Left != nil {
				treearr = append(treearr, treearr[0].Left)
			}

			if treearr[0].Right != nil {
				treearr = append(treearr, treearr[0].Right)
			}

			treearr[0] = nil
			treearr = treearr[1:]
		}

		arr = append(arr, sum/float64(l))
	}

	return arr
}