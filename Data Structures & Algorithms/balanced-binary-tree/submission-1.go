/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	res := true

	var traverse func(root *TreeNode) int

	traverse = func(root *TreeNode) int{
		if root == nil {
			return 0
		}


		left := traverse(root.Left)
		right := traverse(root.Right)

		// fmt.Println(root.Val)
		// fmt.Println(left)
		// fmt.Println(right)

		diff := left - right

		//abs
		if diff < 0 {
			diff = diff * -1
		}

		if diff > 1 {
			res = false
		}

		max := left
		if max < right {
			max = right
		}

		return 1 + max
	}

	_ = traverse(root)

	return res
}
