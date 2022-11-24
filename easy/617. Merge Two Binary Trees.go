/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	result := &TreeNode{}

	// Проверки на наличие дочерних узлов
	if root1 != nil && root2 == nil {
		result.Val = root1.Val
		result.Left = root1.Left
		result.Right = root1.Right
	} else if root1 == nil && root2 != nil {
		result.Val = root2.Val
		result.Left = root2.Left
		result.Right = root2.Right
	} else if root1 != nil && root2 != nil {
		result.Val = root1.Val + root2.Val
		result.Left = mergeTrees(root1.Left, root2.Left)
		result.Right = mergeTrees(root1.Right, root2.Right)
	}

	return result
}