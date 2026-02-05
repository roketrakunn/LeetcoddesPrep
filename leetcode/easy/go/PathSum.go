package go_test

type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}
 
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil { 
		return  false
	}

	//leaf check 
	if root.Left == nil && root.Right == nil { 
		if targetSum == root.Val { 
			return  true
		}
	}
	remaining := targetSum - root.Val
	return hasPathSum(root.Left, remaining) || hasPathSum(root.Right ,remaining)
}

