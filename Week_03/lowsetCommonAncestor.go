/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

func lowsetCommonAncestor(root p,q *TreeNode)*TreeNode  {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowsetCommonAncestor(root.Left,p,q)
	right := lowsetCommonAncestor(root.Right,p,q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}