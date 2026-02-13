package go_test

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	tree1 := []int{}
	tree2 := []int{}

	tree1 = dfs(root1, tree1)
	tree2 = dfs(root2, tree2)

	if len(tree1) != len(tree2) {
		return false
	}

	for i, val := range tree1 {
		if val != tree2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}

	if root.Left == nil && root.Right == nil {
		arr = append(arr, root.Val)
	}

	arr = dfs(root.Left, arr)
	arr = dfs(root.Right, arr)

	return arr
}

