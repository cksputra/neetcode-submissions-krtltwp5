/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func diameterOfBinaryTree(root *TreeNode) int {

	res := 0

	var traverse func(*TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := traverse(root.Left)
		right := traverse(root.Right)

		sum := left + right

		if res < sum {
			res = sum
		}

		max := left
		if max < right {
			max = right
		}

		return 1 + max
	}

	traverse(root)
	return res
}
