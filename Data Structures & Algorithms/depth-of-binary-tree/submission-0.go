/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, counter int) int {
	if root == nil {
		return counter
	}
	counter++
	L := traverse(root.Left, counter)
	R := traverse(root.Right, counter)

	max := L
	if R > max {
		max = R
	}

	return max
}


func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	counter := 0
	return traverse(root, counter)
}
