/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    return helper(root, result)
}

func helper(node *TreeNode, result []int) []int {
    if node == nil {
        return result
    }
    result = helper(node.Left, result)
    result = append(result, node.Val)
    result = helper(node.Right, result)
    return result
}