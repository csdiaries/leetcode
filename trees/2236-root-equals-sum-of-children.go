/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func checkTree(root *TreeNode) bool {
    return root.Val == sum(root.Left)+sum(root.Right)
}

func sum(node *TreeNode) int {
    if node == nil {
        return 0
    }

    leftSum := sum(node.Left)
    rightSum := sum(node.Right)

    return node.Val + leftSum + rightSum
}