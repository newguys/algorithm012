var res[]int
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func preorderTraversal(root *TreeNode) []int  {
	res  = make([]int, 0)
	dfs(root)
	return res
}

func dfs(root *TreeNode){
	if root != nil {
		res = append(res,root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
}


//递归
func preorderTravesal(root *TreeNode) []int {
	res := []int{}
	var stack []*TreeNode

	for 0 <len(stack) || root != nil{
		for root != nil{
			res = append(res,root.Val)
			stack = append(stack,root.Right)
			root = root.Left
		}
		index := len(stack)-1
		root = stack[index]
		stack = stack[:index]
	}
	return res
}
